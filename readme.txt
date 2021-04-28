Command us2stdio dials server serving on path and redirects it's io to stdio.
To mount a server serving on my9pserver.socket in acme-sac:
on windows:
go get github.com/fgergo/us2stdio
in acme-sac:
Local mount { {os us2stdio.exe my9pserver.socket}>[1=0]} /n/my9pserver
