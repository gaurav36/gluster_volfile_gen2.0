# gluster_volfile_gen2.0
Volfile generation API for glusterd 2.0

Setup:

[1]. For downloading volfile gen repo do

the master branch of this repo will have graph generation using link list
https://github.com/gaurav36/gluster_volfile_gen2.0/


if you want to test graph generation using slice then do

go get https://github.com/gaurav36/gluster_volfile_gen2.0/blob/BASE-graph-generation-using-slice



[2]. After that build and install binary file using
 go install ./...

 It will generate binary file in bin directory

[3]. Run the binary file using following argument

if you want to see what all these argument are you can type
./<binary file name> -h

it will show you
Usage of ./gluster_volfile_gen2.0:
  -brick="": whether volfile contain local node brick or all nodes brick
  -daemon="": daemon for which volfile generated
  -volname="": volume name
  -vpath="": volfile path

you can run this binary by
./<binary> -vpath <volfile path> -volname <VOLNAME> -daemon <Daemon name> -brick <locally/globally>

During executing of this binary it will ask for how many brick present for volume <VOLNAME>
you need to pass total number of brick count

after that it will generat volfile in <volfile path>
