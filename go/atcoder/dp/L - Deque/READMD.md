<p>Score : <var>100</var> points</p>

<div class="part">
<section>
<h3>Problem Statement</h3><p>Taro and Jiro will play the following game against each other.</p>
<p>Initially, they are given a sequence <var>a = (a_1, a_2, \ldots, a_N)</var>.
Until <var>a</var> becomes empty, the two players perform the following operation alternately, starting from Taro:</p>
<ul>
<li>Remove the element at the beginning or the end of <var>a</var>. The player earns <var>x</var> points, where <var>x</var> is the removed element.</li>
</ul>
<p>Let <var>X</var> and <var>Y</var> be Taro's and Jiro's total score at the end of the game, respectively.
Taro tries to maximize <var>X - Y</var>, while Jiro tries to minimize <var>X - Y</var>.</p>
<p>Assuming that the two players play optimally, find the resulting value of <var>X - Y</var>.</p>
</section>
</div>

<div class="part">
<section>
<h3>Constraints</h3><ul>
<li>All values in input are integers.</li>
<li><var>1 \leq N \leq 3000</var></li>
<li><var>1 \leq a_i \leq 10^9</var></li>
</ul>
</section>
</div>

<hr />
<div class="io-style">
<div class="part">
<section>
<h3>Input</h3><p>Input is given from Standard Input in the following format:</p>
<pre><var>N</var>
<var>a_1</var> <var>a_2</var> <var>\ldots</var> <var>a_N</var>
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Output</h3><p>Print the resulting value of <var>X - Y</var>, assuming that the two players play optimally.</p>
</section>
</div>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 1</h3><pre>4
10 80 90 30
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 1</h3><pre>10
</pre>

<p>The game proceeds as follows when the two players play optimally (the element being removed is written bold):</p>
<ul>
<li>Taro: (10, 80, 90, <strong>30</strong>) → (10, 80, 90)</li>
<li>Jiro: (10, 80, <strong>90</strong>) → (10, 80)</li>
<li>Taro: (10, <strong>80</strong>) → (10)</li>
<li>Jiro: (<strong>10</strong>) → ()</li>
</ul>
<p>Here, <var>X = 30 + 80 = 110</var> and <var>Y = 90 + 10 = 100</var>.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 2</h3><pre>3
10 100 10
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 2</h3><pre>-80
</pre>

<p>The game proceeds, for example, as follows when the two players play optimally:</p>
<ul>
<li>Taro: (<strong>10</strong>, 100, 10) → (100, 10)</li>
<li>Jiro: (<strong>100</strong>, 10) → (10)</li>
<li>Taro: (<strong>10</strong>) → ()</li>
</ul>
<p>Here, <var>X = 10 + 10 = 20</var> and <var>Y = 100</var>.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 3</h3><pre>1
10
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 3</h3><pre>10
</pre>

</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 4</h3><pre>10
1000000000 1 1000000000 1 1000000000 1 1000000000 1 1000000000 1
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 4</h3><pre>4999999995
</pre>

<p>The answer may not fit into a 32-bit integer type.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 5</h3><pre>6
4 2 9 7 1 5
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 5</h3><pre>2
</pre>

<p>The game proceeds, for example, as follows when the two players play optimally:</p>
<ul>
<li>Taro: (4, 2, 9, 7, 1, <strong>5</strong>) → (4, 2, 9, 7, 1)</li>
<li>Jiro: (<strong>4</strong>, 2, 9, 7, 1) → (2, 9, 7, 1)</li>
<li>Taro: (2, 9, 7, <strong>1</strong>) → (2, 9, 7)</li>
<li>Jiro: (2, 9, <strong>7</strong>) → (2, 9)</li>
<li>Taro: (2, <strong>9</strong>) → (2)</li>
<li>Jiro: (<strong>2</strong>) → ()</li>
</ul>
<p>Here, <var>X = 5 + 1 + 9 = 15</var> and <var>Y = 4 + 7 + 2 = 13</var>.</p></section>
</div>
