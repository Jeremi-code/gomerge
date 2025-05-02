package app


func SwitchBranch(branchName string) {
	RunCommand("git stash")
	RunCommand("git checkout " + branchName)
}

func PullBranch() {
	RunCommand("git pull")
}

func MergeBranch(sourceBranchName string, targetBranchName string) {
	RunCommand("git switch " + targetBranchName)
	RunCommand("git merge " + sourceBranchName)
}