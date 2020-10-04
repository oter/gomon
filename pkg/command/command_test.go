package command

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

// NOTE: All the tests expect $PROJECT_ROOT variable to be set and pointed to project's root path.

// Test expects test-app-args binary to be built and located in $PROJECT_ROOT/bin folder.
func TestPassedArguments(t *testing.T) {
	projectPath := os.Getenv("PROJECT_ROOT")
	if projectPath == "" {
		t.Fatal("PROJECT_PATH env var expected to be set")
	}

	testCases := []struct {
		CaseName         string
		Arguments        []string
		ExpectedExitCode int
		ReturnsError     bool
	}{
		{
			CaseName:         "wrong argument passed",
			Arguments:        []string{"-wrongArgument"},
			ExpectedExitCode: 127,
			ReturnsError:     true,
		},
		{
			CaseName:         "no arguments",
			Arguments:        []string{},
			ExpectedExitCode: 0,
			ReturnsError:     false,
		},
		{
			CaseName:         "several arguments",
			Arguments:        []string{"-arg1", "-arg1", "-arg2", "-arg3", "-arg3"},
			ExpectedExitCode: 5,
			ReturnsError:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.CaseName, func(t *testing.T) {
			params := &Params{
				Path: path.Join(projectPath, "/bin/test-app-args"),
				Args: tc.Arguments,
			}
			cmd, err := New(params)
			require.Nil(t, err)

			waitError := cmd.Wait()
			if tc.ReturnsError {
				require.Error(t, waitError)
			} else {
				require.NoError(t, waitError)
			}

			exitCode, ok := cmd.ExitCode()
			require.True(t, ok)
			require.Equal(t, tc.ExpectedExitCode, exitCode)
		})
	}
}

// Test expects test-app-kill binary to be built and located in $PROJECT_ROOT/bin folder.
func TestSigkill(t *testing.T) {
	projectPath := os.Getenv("PROJECT_ROOT")
	if projectPath == "" {
		t.Fatal("PROJECT_PATH env var expected to be set")
	}

	t.Run("sigterm", func(t *testing.T) {
		cmd, err := New(&Params{Path: path.Join(projectPath, "/bin/test-app-kill")})
		require.Nil(t, err)

		require.NoError(t, cmd.Sigkill())

		waitError := cmd.Wait()
		require.Error(t, waitError)

		exitCode, ok := cmd.ExitCode()
		require.True(t, ok)
		require.Equal(t, -1, exitCode) // -1 means the child process was successfully killed
	})

	t.Run("timeout", func(t *testing.T) {
		cmd, err := New(&Params{Path: path.Join(projectPath, "/bin/test-app-kill")})
		require.Nil(t, err)

		waitError := cmd.Wait()
		require.Error(t, waitError)

		exitCode, ok := cmd.ExitCode()
		require.True(t, ok)
		require.Equal(t, 128, exitCode)
	})
}
