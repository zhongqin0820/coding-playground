<p>Score : <var>100</var> points</p>

<div class="part">
<section>
<h3>Problem Statement</h3><p>There are <var>N</var> items, numbered <var>1, 2, \ldots, N</var>.
For each <var>i</var> (<var>1 \leq i \leq N</var>), Item <var>i</var> has a weight of <var>w_i</var> and a value of <var>v_i</var>.</p>
<p>Taro has decided to choose some of the <var>N</var> items and carry them home in a knapsack.
The capacity of the knapsack is <var>W</var>, which means that the sum of the weights of items taken must be at most <var>W</var>.</p>
<p>Find the maximum possible sum of the values of items that Taro takes home.</p>
</section>
</div>

<div class="part">
<section>
<h3>Constraints</h3><ul>
<li>All values in input are integers.</li>
<li><var>1 \leq N \leq 100</var></li>
<li><var>1 \leq W \leq 10^9</var></li>
<li><var>1 \leq w_i \leq W</var></li>
<li><var>1 \leq v_i \leq 10^3</var></li>
</ul>
</section>
</div>

<hr />
<div class="io-style">
<div class="part">
<section>
<h3>Input</h3><p>Input is given from Standard Input in the following format:</p>
<pre><var>N</var> <var>W</var>
<var>w_1</var> <var>v_1</var>
<var>w_2</var> <var>v_2</var>
<var>:</var>
<var>w_N</var> <var>v_N</var>
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Output</h3><p>Print the maximum possible sum of the values of items that Taro takes home.</p>
</section>
</div>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 1</h3><pre>3 8
3 30
4 50
5 60
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 1</h3><pre>90
</pre>

<p>Items <var>1</var> and <var>3</var> should be taken.
Then, the sum of the weights is <var>3 + 5 = 8</var>, and the sum of the values is <var>30 + 60 = 90</var>.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 2</h3><pre>1 1000000000
1000000000 10
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 2</h3><pre>10
</pre>

</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 3</h3><pre>6 15
6 5
5 6
6 4
6 6
3 5
7 2
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 3</h3><pre>17
</pre>

<p>Items <var>2, 4</var> and <var>5</var> should be taken.
Then, the sum of the weights is <var>5 + 6 + 3 = 14</var>, and the sum of the values is <var>6 + 6 + 5 = 17</var>.</p></section>
</div>
