<p>Score : <var>100</var> points</p>

<div class="part">
<section>
<h3>Problem Statement</h3><p>There are <var>N</var> dishes, numbered <var>1, 2, \ldots, N</var>.
Initially, for each <var>i</var> (<var>1 \leq i \leq N</var>), Dish <var>i</var> has <var>a_i</var> (<var>1 \leq a_i \leq 3</var>) pieces of sushi on it.</p>
<p>Taro will perform the following operation repeatedly until all the pieces of sushi are eaten:</p>
<ul>
<li>Roll a die that shows the numbers <var>1, 2, \ldots, N</var> with equal probabilities, and let <var>i</var> be the outcome. If there are some pieces of sushi on Dish <var>i</var>, eat one of them; if there is none, do nothing.</li>
</ul>
<p>Find the expected number of times the operation is performed before all the pieces of sushi are eaten.</p>
</section>
</div>

<div class="part">
<section>
<h3>Constraints</h3><ul>
<li>All values in input are integers.</li>
<li><var>1 \leq N \leq 300</var></li>
<li><var>1 \leq a_i \leq 3</var></li>
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
<h3>Output</h3><p>Print the expected number of times the operation is performed before all the pieces of sushi are eaten.
The output is considered correct when the relative difference is not greater than <var>10^{-9}</var>.</p>
</section>
</div>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 1</h3><pre>3
1 1 1
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 1</h3><pre>5.5
</pre>

<p>The expected number of operations before the first piece of sushi is eaten, is <var>1</var>.
After that, the expected number of operations before the second sushi is eaten, is <var>1.5</var>.
After that, the expected number of operations before the third sushi is eaten, is <var>3</var>.
Thus, the expected total number of operations is <var>1 + 1.5 + 3 = 5.5</var>.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 2</h3><pre>1
3
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 2</h3><pre>3
</pre>

<p>Outputs such as <code>3.00</code>, <code>3.000000003</code> and <code>2.999999997</code> will also be accepted.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 3</h3><pre>2
1 2
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 3</h3><pre>4.5
</pre>

</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 4</h3><pre>10
1 3 2 3 3 2 3 2 1 3
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 4</h3><pre>54.48064457488221
</pre></section>
</div>
