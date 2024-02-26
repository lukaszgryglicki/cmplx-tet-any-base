#!/bin/bash
WIDTH=800 HEIGHT=600 DEBUG='' go run natural_tetration.go 1.4 0.1 -2 2 -2 2 && convert R.png G.png B.png -combine A.png && ls -l A.png
