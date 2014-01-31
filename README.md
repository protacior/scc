scc
===

Website for SCC

Tools:
  Hosted on Github
  Integrate with Continuous Integration Service like Drone or Werker

Front End:
  HTML 5
  CSS3
  Javascript
  jQuery

Back End:
  Go
  
Web Host:
  Google App Engine

Resources for learning Go: Expected time 2-5 hours
  Installation:
    If mac use "brew install go" (Homebrew)
    If linux use package manager
  http://golang.org/doc/
    Getting Started
    A Tour of Go
    How to write Go code
    Effective Go
  http://golang.org/doc/articles/wiki/
    Build a basic webservice

Resources for integrating with App Engine: 
  https://developers.google.com/appengine/docs/go/gettingstarted/introduction
    Go through the entire tutorial
  https://developers.google.com/appengine/docs/go/datastore/
    Go through how to store data (these roommate request posts)

Major Sections:
User Sign Up
  Series of "slides"
  Two options: Roommate Request or Find Roommates in Atea
  Roommate Request:
    Name/Password -- HTTPS -- could do google authentication (will ask)
    Email
    Start Date-End Date
    Price Range - Sliding Scale
    City, State -- link with google maps to verify and put in DB in standard form
    Gender, and if a certain gender is required for roommates
    What I want in a roommate?
    What I am as a roommate?
    Submit --> User Verification
  Find Roommates:
    Name/Password -- HTTPS -- could do google authentication (will ask)
    Email
    Sign up for Email Notifications [weekly digest?]
    Submit --> User Verification
  

User Verification
  For sign up, send url with unique link (SHA hash)
  Make verification page
  Update DB entry

Search for Roommates:
  Filter results based off area/distance -- link with Google Maps?
  Filter based off city, state
  Filter based off price range
  Sort based off distance from...
  
Email Notifications:
  Option to sign up for email notifications
  Generate email notifications in a weekly digest (for certain areas)
    Send them to the people who have signed up for Email Notifications
    
Database:
User: First Name, Last Name, Email, Password, Verified
Request: First Name, Last Name, Email, Start Date, End Date, Low Price, High Price, City, State, Gender, What I want, What I am

Specific Entries:
User:
  First+Last Name: Normalized Form 
  Email: Valid "yale" email address
  Password: Encrypted
  Verified: Boolean

Request:
  First+Last Name: Normalized Form
  Email: Any email address
  Start-Date: Valid Start Date
  End-Date: Valid End Date (or a negative for undetermined)
  Low Price: Integer
  High Price: Integer
  City: Valid City -- Verified by google maps
  State: Valid State -- Verified by gm api
  Gender: Bool (True male, False female)
  What I want: Text (Impose limit of 250 characters)
  What I am: Text (Impose limit of 250 characters)
