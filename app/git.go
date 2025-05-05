package app


func SwitchBranch(branchName string) {
	RunCommand("git checkout " + branchName)
}

func PullBranch() {
	RunCommand("git stash")
	RunCommand("git pull")
}

func MergeBranch(sourceBranchName string, targetBranchName string) {
	SwitchBranch(targetBranchName)
	RunCommand("git merge " + sourceBranchName)
	RunCommand("git stash pop")
}