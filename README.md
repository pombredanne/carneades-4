
These are the source files of Version 4 of the
Carneades argumentation system, written in the Go programming language.

This source code is subject to the terms of the Mozilla Public
License, version 2.0 (MPL-2.0). If a copy of the MPL was not
distributed with this software, it is also available online at
<http://mozilla.org/MPL/2.0/>.  For futher information about the MPL see <http://www.mozilla.org/MPL/2.0/FAQ.html>.

This version of Carneades consists of:

- An evaluator for structured arguments, based on a new version of the 
Carneades Argument Evaluation Structures (CAES) formal model of argument. 
New in this version of CAES is support for cyclic argument graphs, cumulative
arguments, issue-based information systems (IBIS) and multicriteria decision analysis.
Argument graphs can be represented in YAML, AIF, LKIF, and CAF and exported to DOT, GraphML and YAML.
- An implementation of a solver for Dung abstract argumentation frameworks,
using grounded, complete, preferred and stable semantics. Argumentation Frameworks
can be represented using the [Trivial Graph Format](https://en.wikipedia.org/wiki/Trivial_Graph_Format). The computed extensions can be exported to DOT, GraphML and plain text.

Carneades can visualize both Dung abstract argumentation frameworks and CAES argument graphs using [DOT](http://www.graphviz.org/content/dot-language) and [GraphML](https://en.wikipedia.org/wiki/GraphML).  We recommend using the free [yEd](http://www.yworks.com/yed) GraphML editor to view the GraphML files.

Please see the Carneades blog at <https://carneades.github.io/> for
announcements about the further development of this version.

See the INSTALL.md file in the same directory as this README for
installation instructions.

