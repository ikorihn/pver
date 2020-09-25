package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/tcnksm/go-gitconfig"
)

type repository struct {
	wd string
}

func NewRepository(wd string) *repository {
	return &repository{wd: wd}
}

func (r *repository) CommitUpdate(filePath string, updateVer string) error {
	repo, err := git.PlainOpen(r.wd)
	if err != nil {
		if errors.Is(err, git.ErrRepositoryNotExists) {
			// make error message git-like
			fmt.Println("fatal: not a git repository")
			os.Exit(1)
		}
		return nil
	}

	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	_, err = w.Add(filePath)
	if err != nil {
		return err
	}

	email, err := gitconfig.Email()
	if err != nil {
		return err
	}
	name, err := gitconfig.Username()
	if err != nil {
		return err
	}

	commitMessage := fmt.Sprintf("version up to %s", updateVer)
	fmt.Printf("commit: %s", commitMessage)
	w.Commit(commitMessage, &git.CommitOptions{
		Author: &object.Signature{
			Email: email,
			Name:  name,
			When:  time.Now(),
		},
	})

	return nil
}

func (r *repository) CreateBranch(name string) error {
	repo, err := git.PlainOpen(r.wd)
	if err != nil {
		return err
	}

	headRef, err := repo.Head()
	if err != nil {
		return err
	}

	branchName := plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", name))
	ref := plumbing.NewHashReference(branchName, headRef.Hash())

	return repo.Storer.SetReference(ref)
}
