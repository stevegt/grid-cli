#!/bin/bash

# padsp signalgen -t 100m sin 444

chatfile=/tmp/$$.chat

inputs=$(ls *.md)
inputs=$(echo $inputs | sed 's/ /,/g')

sysmsg="You are an expert technical writer and software architect. 
    You are writing a book about PromiseGrid.  Each of the input files
    is a chapter in the book.  "
prompt="Examine the chapters.  Look for conflicts, redundancies, and
    inconsistencies.  Decide which chapter most needs revision.  Write the
    filename of the chapter as your answer.  Your answer must be only 
    the filename of the chapter that most needs revision.  Add no
    other information to your answer."
# cmd="grok chat $chatfile -N -s $sysmsg -i $inputs -o v2/doc/target-filename"
echo $cmd
echo $prompt | grok chat $chatfile -N -s "$sysmsg" -i $inputs


