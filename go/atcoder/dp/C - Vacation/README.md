<p>Score : <var>100</var> points</p>

<div class="part">
<section>
<h3>Problem Statement</h3><p>Taro's summer vacation starts tomorrow, and he has decided to make plans for it now.</p>
<p>The vacation consists of <var>N</var> days.
For each <var>i</var> (<var>1 \leq i \leq N</var>), Taro will choose one of the following activities and do it on the <var>i</var>-th day:</p>
<ul>
<li>A: Swim in the sea. Gain <var>a_i</var> points of happiness.</li>
<li>B: Catch bugs in the mountains. Gain <var>b_i</var> points of happiness.</li>
<li>C: Do homework at home. Gain <var>c_i</var> points of happiness.</li>
</ul>
<p>As Taro gets bored easily, he cannot do the same activities for two or more consecutive days.</p>
<p>Find the maximum possible total points of happiness that Taro gains.</p>
</section>
</div>

<div class="part">
<section>
<h3>Constraints</h3><ul>
<li>All values in input are integers.</li>
<li><var>1 \leq N \leq 10^5</var></li>
<li><var>1 \leq a_i, b_i, c_i \leq 10^4</var></li>
</ul>
</section>
</div>

<hr />
<div class="io-style">
<div class="part">
<section>
<h3>Input</h3><p>Input is given from Standard Input in the following format:</p>
<pre><var>N</var>
<var>a_1</var> <var>b_1</var> <var>c_1</var>
<var>a_2</var> <var>b_2</var> <var>c_2</var>
<var>:</var>
<var>a_N</var> <var>b_N</var> <var>c_N</var>
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Output</h3><p>Print the maximum possible total points of happiness that Taro gains.</p>
</section>
</div>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 1</h3><pre>3
10 40 70
20 50 80
30 60 90
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 1</h3><pre>210
</pre>

<p>If Taro does activities in the order C, B, C, he will gain <var>70 + 50 + 90 = 210</var> points of happiness.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 2</h3><pre>1
100 10 1
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 2</h3><pre>100
</pre>

</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 3</h3><pre>7
6 7 8
8 8 3
2 5 2
7 8 6
4 6 8
2 3 4
7 5 1
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 3</h3><pre>46
</pre>

<p>Taro should do activities in the order C, A, B, A, C, B, A.</p></section>
</div>
