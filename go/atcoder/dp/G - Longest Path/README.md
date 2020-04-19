<p>Score : <var>100</var> points</p>

<div class="part">
<section>
<h3>Problem Statement</h3><p>There is a directed graph <var>G</var> with <var>N</var> vertices and <var>M</var> edges.
The vertices are numbered <var>1, 2, \ldots, N</var>, and for each <var>i</var> (<var>1 \leq i \leq M</var>), the <var>i</var>-th directed edge goes from Vertex <var>x_i</var> to <var>y_i</var>.
<var>G</var> <strong>does not contain directed cycles</strong>.</p>
<p>Find the length of the longest directed path in <var>G</var>.
Here, the length of a directed path is the number of edges in it.</p>
</section>
</div>

<div class="part">
<section>
<h3>Constraints</h3><ul>
<li>All values in input are integers.</li>
<li><var>2 \leq N \leq 10^5</var></li>
<li><var>1 \leq M \leq 10^5</var></li>
<li><var>1 \leq x_i, y_i \leq N</var></li>
<li>All pairs <var>(x_i, y_i)</var> are distinct.</li>
<li><var>G</var> <strong>does not contain directed cycles</strong>.</li>
</ul>
</section>
</div>

<hr />
<div class="io-style">
<div class="part">
<section>
<h3>Input</h3><p>Input is given from Standard Input in the following format:</p>
<pre><var>N</var> <var>M</var>
<var>x_1</var> <var>y_1</var>
<var>x_2</var> <var>y_2</var>
<var>:</var>
<var>x_M</var> <var>y_M</var>
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Output</h3><p>Print the length of the longest directed path in <var>G</var>.</p>
</section>
</div>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 1</h3><pre>4 5
1 2
1 3
3 2
2 4
3 4
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 1</h3><pre>3
</pre>

<p>The red directed path in the following figure is the longest:</p>
<p><img alt="" src="https://img.atcoder.jp/dp/longest_0_muffet.png" /></p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 2</h3><pre>6 3
2 3
4 5
5 6
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 2</h3><pre>2
</pre>

<p>The red directed path in the following figure is the longest:</p>
<p><img alt="" src="https://img.atcoder.jp/dp/longest_1_muffet.png" /></p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 3</h3><pre>5 8
5 3
2 3
2 4
5 2
5 1
1 4
4 3
1 3
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 3</h3><pre>3
</pre>

<p>The red directed path in the following figure is one of the longest:</p>
<p><img alt="" src="https://img.atcoder.jp/dp/longest_2_muffet.png" /></p></section>
</div>
