# Cant

`cant` - _language peculiar to a specified group or profession and regarded with disparagement._ (One of many definitions found on the Internet. [See](https://en.wikipedia.org/wiki/Cant_(language)) [also...](https://en.wikipedia.org/wiki/Thieves%27_cant))

Cant is a reimagination of [Apache Ant](https://ant.apache.org/manual/index.html).  Ant is, in my opinion, a beautiful hot mess of a language.  It's the [Rodney Dangerfield](https://en.wikipedia.org/wiki/Rodney_Dangerfield) of [multi-paradigm](https://en.wikipedia.org/wiki/Comparison_of_multi-paradigm_programming_languages) programming languages.

### Paradigms in Ant
* __\<target\>__ - dependency satisfaction, like old-school Unix `make`
* __\<property\>__ - global write-once properties ...
* __\<local\>__ - ... except when
* __\<macrodef\>__ - procedural/imperative
* __\<taskdef\>__ - host language extensions
* __\<fileset\>__ - declarative
* __\<parallel\>__ - concurrency
* __\<and\>\<or\>\<not\>__ - logic
* __[AntContrib](http://ant-contrib.sourceforge.net/tasks/tasks/index.html)__ - with `<if>`, `<foreach>`, etc...
* regular expressions, filename patterns
* embedded in XML

