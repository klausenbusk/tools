package util

import (
	"context"
	"os/exec"
	"strings"
)

// UpdateGitRepo git pulls and rebases the repository.
func UpdateGitRepo(ctx context.Context, gitpath string) {
	args := []string{"pull", "--rebase"}

	cmd := exec.Command("git", args...)
	cmd.Dir = gitpath
	Debugf(ctx, "%s: run %s\n", gitpath, cmd)
	CheckCmd(cmd.CombinedOutput())
}

func FileExistsInGit(ctx context.Context, gitpath string, filepath string) bool {
	args := []string{"ls-files", "--error-unmatch", filepath}

	cmd := exec.Command("git", args...)
	cmd.Dir = gitpath
	Debugf(ctx, "%s: run %s\n", gitpath, cmd)

	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func FileIgnoredByGit(ctx context.Context, gitpath string, filepath string) bool {
	args := []string{"check-ignore", "--quiet", "--no-index", strings.TrimLeft(filepath, gitpath)}

	cmd := exec.Command("git", args...)
	cmd.Dir = gitpath
	Debugf(ctx, "%s: run %s\n", gitpath, cmd)

	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
