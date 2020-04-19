<p>Score : <var>100</var> points</p>

<div class="part">
<section>
<h3>Problem Statement</h3><p>There are <var>N</var> children, numbered <var>1, 2, \ldots, N</var>.</p>
<p>They have decided to share <var>K</var> candies among themselves.
Here, for each <var>i</var> (<var>1 \leq i \leq N</var>), Child <var>i</var> must receive between <var>0</var> and <var>a_i</var> candies (inclusive).
Also, no candies should be left over.</p>
<p>Find the number of ways for them to share candies, modulo <var>10^9 + 7</var>.
Here, two ways are said to be different when there exists a child who receives a different number of candies.</p>
</section>
</div>

<div class="part">
<section>
<h3>Constraints</h3><ul>
<li>All values in input are integers.</li>
<li><var>1 \leq N \leq 100</var></li>
<li><var>0 \leq K \leq 10^5</var></li>
<li><var>0 \leq a_i \leq K</var></li>
</ul>
</section>
</div>

<hr />
<div class="io-style">
<div class="part">
<section>
<h3>Input</h3><p>Input is given from Standard Input in the following format:</p>
<pre><var>N</var> <var>K</var>
<var>a_1</var> <var>a_2</var> <var>\ldots</var> <var>a_N</var>
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Output</h3><p>Print the number of ways for the children to share candies, modulo <var>10^9 + 7</var>.</p>
</section>
</div>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 1</h3><pre>3 4
1 2 3
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 1</h3><pre>5
</pre>

<p>There are five ways for the children to share candies, as follows:</p>
<ul>
<li><var>(0, 1, 3)</var></li>
<li><var>(0, 2, 2)</var></li>
<li><var>(1, 0, 3)</var></li>
<li><var>(1, 1, 2)</var></li>
<li><var>(1, 2, 1)</var></li>
</ul>
<p>Here, in each sequence, the <var>i</var>-th element represents the number of candies that Child <var>i</var> receives.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 2</h3><pre>1 10
9
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 2</h3><pre>0
</pre>

<p>There may be no ways for the children to share candies.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 3</h3><pre>2 0
0 0
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 3</h3><pre>1
</pre>

<p>There is one way for the children to share candies, as follows:</p>
<ul>
<li><var>(0, 0)</var></li>
</ul>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 4</h3><pre>4 100000
100000 100000 100000 100000
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 4</h3><pre>665683269
</pre>

<p>Be sure to print the answer modulo <var>10^9 + 7</var>.</p></section>
</div>
