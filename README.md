golang-source
=============


    All source codes come from youngcy.


    While learning the go programming language, I write this codes. As a programming language,

not only thinking but also exercising. The go programming language's syntax can be found in

my gobook repository. This is also my reference book, all I had studied is coms from here.


    First, you should think about why you begin to study this language.


    For me, there is no doubt that I am a google fance. Since my first time use google product-chrome

in 2012, I thought I was found of it. Then as the experience enriching, more and more google producs

appeared in my life. It is said that the google company is the best work place for ITs. So I dream that

I can work in the google company someday. This is my indivadual reson. Now, let's begin. In fact, thaer

are many resons force me to learn it. The most important afact is go language is a language for web 

server(If you have different views, don't argue with me.). And it is easy and quik to learn. The biggest

character is favor the goruntine.


    Second, give up the C language habit and you must do it.
    
    At first time I using go language, I really don't fit it. The go syntax is different from C language.
But if you are a C programmer, you will crazy while you first time learn the go language. You can refer the
gobook repository. But if you had learnt it, you will love it as me.


    Third, go language environment setup.
    In ubuntu, we can use:
        sudo apt-get install golang-go
    (Note: the go comes from google, the google use linux for their work environment.)
    
    
    You should setup the GOPATH in correct method. Here I gives a example path to show you how to use GOPATH
    --<sorter>
      --<src>
        --<algorithms>
          --quiksort.go
          --bubblesort.go
        --main.go
      --<bin>
      --<pkg>
      
      Then we set the GOPATH as the current path of sorter. We presume that the sorter path is ~/golang/sorter
We use command line:
    export GOPATH=~/golang/sorter
    source ~/.bashrc
Then you can use go build or go install to compile to the pkg and bin directory. More infromation please refer to gobook.

    OK, wish you have fun in learning go language.
