<p>Score : <var>100</var> points</p>

<div class="part">
<section>
<h3>Problem Statement</h3><p>Let <var>N</var> be a positive odd number.</p>
<p>There are <var>N</var> coins, numbered <var>1, 2, \ldots, N</var>.
For each <var>i</var> (<var>1 \leq i \leq N</var>), when Coin <var>i</var> is tossed, it comes up heads with probability <var>p_i</var> and tails with probability <var>1 - p_i</var>.</p>
<p>Taro has tossed all the <var>N</var> coins.
Find the probability of having more heads than tails.</p>
</section>
</div>

<div class="part">
<section>
<h3>Constraints</h3><ul>
<li><var>N</var> is an odd number.</li>
<li><var>1 \leq N \leq 2999</var></li>
<li><var>p_i</var> is a real number and has two decimal places.</li>
<li><var>0 &lt; p_i &lt; 1</var></li>
</ul>
</section>
</div>

<hr />
<div class="io-style">
<div class="part">
<section>
<h3>Input</h3><p>Input is given from Standard Input in the following format:</p>
<pre><var>N</var>
<var>p_1</var> <var>p_2</var> <var>\ldots</var> <var>p_N</var>
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Output</h3><p>Print the probability of having more heads than tails.
The output is considered correct when the absolute error is not greater than <var>10^{-9}</var>.</p>
</section>
</div>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 1</h3><pre>3
0.30 0.60 0.80
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 1</h3><pre>0.612
</pre>

<p>The probability of each case where we have more heads than tails is as follows:</p>
<ul>
<li>The probability of having <var>(Coin 1, Coin 2, Coin 3) = (Head, Head, Head)</var> is <var>0.3 × 0.6 × 0.8 = 0.144</var>;</li>
<li>The probability of having <var>(Coin 1, Coin 2, Coin 3) = (Tail, Head, Head)</var> is <var>0.7 × 0.6 × 0.8 = 0.336</var>;</li>
<li>The probability of having <var>(Coin 1, Coin 2, Coin 3) = (Head, Tail, Head)</var> is <var>0.3 × 0.4 × 0.8 = 0.096</var>;</li>
<li>The probability of having <var>(Coin 1, Coin 2, Coin 3) = (Head, Head, Tail)</var> is <var>0.3 × 0.6 × 0.2 = 0.036</var>.</li>
</ul>
<p>Thus, the probability of having more heads than tails is <var>0.144 + 0.336 + 0.096 + 0.036 = 0.612</var>.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 2</h3><pre>1
0.50
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 2</h3><pre>0.5
</pre>

<p>Outputs such as <code>0.500</code>, <code>0.500000001</code> and <code>0.499999999</code> are also considered correct.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 3</h3><pre>5
0.42 0.01 0.42 0.99 0.42
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 3</h3><pre>0.3821815872
</pre></section>
</div>
