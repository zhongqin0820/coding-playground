<p>Score : <var>100</var> points</p>

<div class="part">
<section>
<h3>Problem Statement</h3><p>There is a grid with <var>H</var> horizontal rows and <var>W</var> vertical columns.
Let <var>(i, j)</var> denote the square at the <var>i</var>-th row from the top and the <var>j</var>-th column from the left.</p>
<p>For each <var>i</var> and <var>j</var> (<var>1 \leq i \leq H</var>, <var>1 \leq j \leq W</var>), Square <var>(i, j)</var> is described by a character <var>a_{i, j}</var>.
If <var>a_{i, j}</var> is <code>.</code>, Square <var>(i, j)</var> is an empty square; if <var>a_{i, j}</var> is <code>#</code>, Square <var>(i, j)</var> is a wall square.
It is guaranteed that Squares <var>(1, 1)</var> and <var>(H, W)</var> are empty squares.</p>
<p>Taro will start from Square <var>(1, 1)</var> and reach <var>(H, W)</var> by repeatedly moving right or down to an adjacent empty square.</p>
<p>Find the number of Taro's paths from Square <var>(1, 1)</var> to <var>(H, W)</var>.
As the answer can be extremely large, find the count modulo <var>10^9 + 7</var>.</p>
</section>
</div>

<div class="part">
<section>
<h3>Constraints</h3><ul>
<li><var>H</var> and <var>W</var> are integers.</li>
<li><var>2 \leq H, W \leq 1000</var></li>
<li><var>a_{i, j}</var> is <code>.</code> or <code>#</code>.</li>
<li>Squares <var>(1, 1)</var> and <var>(H, W)</var> are empty squares.</li>
</ul>
</section>
</div>

<hr />
<div class="io-style">
<div class="part">
<section>
<h3>Input</h3><p>Input is given from Standard Input in the following format:</p>
<pre><var>H</var> <var>W</var>
<var>a_{1, 1}</var><var>\ldots</var><var>a_{1, W}</var>
<var>:</var>
<var>a_{H, 1}</var><var>\ldots</var><var>a_{H, W}</var>
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Output</h3><p>Print the number of Taro's paths from Square <var>(1, 1)</var> to <var>(H, W)</var>, modulo <var>10^9 + 7</var>.</p>
</section>
</div>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 1</h3><pre>3 4
...#
.#..
....
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 1</h3><pre>3
</pre>

<p>There are three paths as follows:</p>
<p><img alt="" src="https://img.atcoder.jp/dp/grid_0_0_muffet.png" /></p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 2</h3><pre>5 2
..
#.
..
.#
..
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 2</h3><pre>0
</pre>

<p>There may be no paths.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 3</h3><pre>5 5
..#..
.....
#...#
.....
..#..
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 3</h3><pre>24
</pre>

</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 4</h3><pre>20 20
....................
....................
....................
....................
....................
....................
....................
....................
....................
....................
....................
....................
....................
....................
....................
....................
....................
....................
....................
....................
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 4</h3><pre>345263555
</pre>

<p>Be sure to print the count modulo <var>10^9 + 7</var>.</p></section>
</div>
