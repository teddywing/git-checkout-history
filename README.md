git-checkout-history
====================

Store history of your previously checked out git branches and quickly go back to recent branches.

Here's an example of how it's used:

	$ git branch
	* master
	
	$ git checkout-store -b new-feature
	Switched to a new branch 'new-feature'
	
	$ git checkout-store -b feature-to-code-review
	Switched to a new branch 'feature-to-code-review'
	
	$ git checkout-store master
	Switched to branch 'master'
	
	$ git checkout-history
	[1] feature-to-code-review
	[2] new-feature
	
	$ git checkout-history 2
	Switched to branch 'new-feature'
	
	$ git checkout-history
	[1] master
	[2] feature-to-code-review
	[3] new-feature


Branch history is stored in a file called `~/.git-checkout-history`. This is a YAML file that contains a list of previously stored branches.

To make the tool easier to use, git aliases can be used:

	$ git config --global alias.chs checkout-store
	$ git config --global alias.ch checkout-history
	
	$ git chs a-branch


## Installation
Visit the [releases](https://github.com/teddywing/git-checkout-history/releases) page and download `git-checkout-history` and `git-checkout-store` for your platform. Put these binaries on your PATH and you should be able to run them using `git`.

### Installing From Source
Run these commands to build `git-checkout-history` and `git-checkout-store`:

	$ go get github.com/teddywing/git-checkout-history/...
	$ go install github.com/teddywing/git-checkout-history/git-checkout-history
	$ go install github.com/teddywing/git-checkout-history/git-checkout-store


## Shell Completion
If you use bash or zsh completion for git, a patch file is included to add this functionality for `git-checkout-store`. The patch enables `checkout-store` to have the same completion functionality as `checkout`. It can be applied as usual:

	$ patch path/to/git-completion.bash < path/to/git-completion.bash.patch


## Known Issues
* History never gets cleared. Not sure if this is actually an issue, but personally I don't really care about branch history from before the current day.


## License
Licensed under the MIT license. See the included LICENSE file.
