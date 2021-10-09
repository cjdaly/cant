# cant

`cant` - _language peculiar to a specified group or profession and regarded with disparagement._ (One of many definitions found on the Internet. [See](https://en.wikipedia.org/wiki/Cant_(language)) [also...](https://en.wikipedia.org/wiki/Thieves%27_cant))

Cant is a reimagination of [Apache Ant](https://ant.apache.org/manual/index.html).  Ant is, in my opinion, a beautiful hot mess of a language.  I think it should be recognized as a [multi-paradigm](https://en.wikipedia.org/wiki/Comparison_of_multi-paradigm_programming_languages) programming language.  Consider:

* __\<target\>__ - dependency satisfaction, like old-school Unix `make`
* __\<property\>__ - global write-once properties ...
* __\<local\>__ - ... except when
* __\<macrodef\>__ - procedural/imperative
* __\<fileset\>__ - declarative
* __\<parallel\>__ - concurrency
* __\<and\>\<or\>\<not\>__ - logic
* __AntContrib__ - with `<if>`, `<foreach>`, etc...
* regular expressions, filename patterns
* embedded in XML

My favorite aspect of Ant is the interplay between `<macrodef>` and the concrete XML syntax.  This is the way to do executable XML!

But there are a lot of loose ends and rough edges.  I want Cant to be like Ant, but cleaned up like it was designed as a multi-paradigm programming language instead of an aggregative hot mess build tool.
