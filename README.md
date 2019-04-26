# Wrk 
Wrk is an open source cli log tool designed for the developer's terminal lifestyle. To contribute submit a pull-request, issue, or feature request.

## Getting started
To get wrking simply go to the [release](https://github.com/juanaveragedeveloper/wrk/releases) page and download the latest release for your OS. Then move it to the directory of where you'll be working at. Example 

`mv wrkcli_windows_amd64.exe /work`

Also, it might help to rename the executable.

`mv wrkcli_windows_amd64.exe wrk.exe`


Then it is recommended to edit the `PATH` system environment variable to have 

`/the/path/to/where/_wrk.exe_is_stored_in/wrk.exe`

## Setup and basic commands
Now, 

`> wrk setup`

will setup up the barebones for wrking. 

Create a notebook on an active project.

`> wrk nb -n "name_of_notebook"`

Now, log messages throughout the day. 

**N.b.** It is important to be in the directory where `wrk setup` was executed otherwise things will not work, for now... look forward to more wrk releases. 

`> wrk log -m "Researched kubernetes best practices" -t "research"` 

By adding a tag with '-t' it makes it easier to search for things through the notebook from the cli.

`> wrk log -f -m "kubernetes" -t "research" ` 
`> "Researched kubernetes best practices `

It also helps to log those times when sidetracked. 

`> wrk log -m "Created SWAPPS/CI github for Chris's azure jenkins pipeline" -t "task"` 

That way it's easy to keep track of all the things that are done throughout the wrk day. 

To switch projects

`> wrk nb -s "name_of_other_notebook" ` 

and start logging away 

`> wrk log -m "Production incident monitoring not working" -t "incident"`

# Future 
There are a lot of ways wrk can be improved. Here are some that the team has thought of so far 

#### Wishlist
<ol>
  <li>Vim / terminal text editor of choice to input long messages.</li>
  <li>Storage of logs in a DB either or local <up to user to configure / login></li>
  <li>Wrking outside of the 'setup' directory. </li>
  <li>Filtering of logs based upon 'time' and 'tag(s)'</li>
  <li>Task start/end time `wrk task -n "debugging" -st' , ` `wrk task -n "debugging" -stp` </li>
</ol>








