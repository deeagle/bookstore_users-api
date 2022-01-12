# Notes

This training seems very confusing, so I take some notes/comments to the sessions.

## Chapter2 
- Lesson 6
  Hey, I totally struggled in the 1st session of the training (lesson 6)!
  
  Why?
  You import your own github packages. That's fine for your local project (even though `go mod init` isn't run?).
  In my environment my IDE (VSCodium) downloads the package from github and I got all the further dependencies for a simple "hello world" called "pong".
  Of course, it searches for a Database that didn't exists.
  	
  I rebuild it without the github package term: https://github.com/deeagle/bookstore_users-api/tree/2.6
  
  For my inhouse courses I run the course on a clean virtual machine. 
  So I find the most of my dependency problems.
  
  Axo, Press "Run" in the jetbrains IDE. Cmon make it more common with `go run main.go` for all without this dependency.
	