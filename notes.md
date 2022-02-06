# Notes

This training seems very confusing, so I take some notes/comments to the sessions.

## Chapter2 
- **6. Basic application structure: MVC pattern** - Code: [2.6]  
  Hey, I totally struggled in the 1st session of the training (lesson 6)!
  
  Why?
  You import your own github packages. That's fine for your local project (even though `go mod init` isn't run?).
  In my environment my IDE (VSCodium) downloads the package from github and I got all the further dependencies for a simple "hello world" called "pong".
  Of course, it searches for a Database that didn't exists.
  	
  I rebuild it without the github package term: https://github.com/deeagle/bookstore_users-api/tree/2.6
  
  For my inhouse courses I run the course on a clean virtual machine. 
  So I find the most of my dependency problems.
  
  Axo, Press "Run" in the jetbrains IDE. Cmon make it more common with `go run main.go` for all without this dependency.

- **7. Starting from our domain** - Code: [2.7]  
  What a chaotic session!  
  The trainer jumps nervous clicky throught the session. Uses different error handling within one block to save one variable (because he replaces it later by a summary function call of the framework). Wild story of API and the whole project to explain the use of predefined error items. Strange (for me) explanations of basic GO concept (multiple return values). He closed the session with some API-Calls within the software [Postman](https://www.postman.com/). Of course without naming or explanation of it. 
  Yes, I got some infos but I still hate this trainer by now!

- **8. How to structure our domain persistence** - Code: [2.8]  
  The session starts with another code stage as the last finished one (`CreateUser() ... return nil, nil`).
  But the trainer is a real developer. He did "boom" sounds when the api crushed (like expected).
  He closes the session with some `git` commands. Of course the first time (no `git` usage in the sessions before)

- **9. Working with dates** - Code: [2.9]  
  Good real world infos of working with time and timezones. The trainer changes the class and folder naming within the session (`date_utils` in the `utils` folder? the object calls included this ... ).

## Chapter 3
- **10. Configure MySQL client in Go** - Code: [3.10]
  Ah okay, you do this course so you know how databases work. And I'm so a real professional that I use Oracle MySQL with a gui tool to prepare my database (what a victim). And I'm so very professional that I use `utf8_spanish2_ci` for my international students. Ok he talks about that but why not go further and be a real trainer.  
  Oh nice, he shows us how to use IDE environment vars ... for GoLand and bash. Ok windows is not professional, but volataile exports? Nah, yml configs are so ugly nowadays ...   
  What a hard course for me ...  
  I add some docker compose and database creation scripts


[3.10]: https://github.com/deeagle/bookstore_users-api/releases/tag/3.10 
[2.9]: https://github.com/deeagle/bookstore_users-api/releases/tag/2.9
[2.8]: https://github.com/deeagle/bookstore_users-api/releases/tag/2.8
[2.7]: https://github.com/deeagle/bookstore_users-api/releases/tag/2.7
[2.6]: https://github.com/deeagle/bookstore_users-api/releases/tag/2.6
