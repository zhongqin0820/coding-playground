<p>Score : <var>100</var> points</p>

<div class="part">
<section>
<h3>Problem Statement</h3><p>There are <var>N</var> stones, numbered <var>1, 2, \ldots, N</var>.
For each <var>i</var> (<var>1 \leq i \leq N</var>), the height of Stone <var>i</var> is <var>h_i</var>.</p>
<p>There is a frog who is initially on Stone <var>1</var>.
He will repeat the following action some number of times to reach Stone <var>N</var>:</p>
<ul>
<li>If the frog is currently on Stone <var>i</var>, jump to Stone <var>i + 1</var> or Stone <var>i + 2</var>. Here, a cost of <var>|h_i - h_j|</var> is incurred, where <var>j</var> is the stone to land on.</li>
</ul>
<p>Find the minimum possible total cost incurred before the frog reaches Stone <var>N</var>.</p>
</section>
</div>

<div class="part">
<section>
<h3>Constraints</h3><ul>
<li>All values in input are integers.</li>
<li><var>2 \leq N \leq 10^5</var></li>
<li><var>1 \leq h_i \leq 10^4</var></li>
</ul>
</section>
</div>

<hr />
<div class="io-style">
<div class="part">
<section>
<h3>Input</h3><p>Input is given from Standard Input in the following format:</p>
<pre><var>N</var>
<var>h_1</var> <var>h_2</var> <var>\ldots</var> <var>h_N</var>
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Output</h3><p>Print the minimum possible total cost incurred.</p>
</section>
</div>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 1</h3><pre>4
10 30 40 20
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 1</h3><pre>30
</pre>

<p>If we follow the path <var>1</var> → <var>2</var> → <var>4</var>, the total cost incurred would be <var>|10 - 30| + |30 - 20| = 30</var>.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 2</h3><pre>2
10 10
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 2</h3><pre>0
</pre>

<p>If we follow the path <var>1</var> → <var>2</var>, the total cost incurred would be <var>|10 - 10| = 0</var>.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 3</h3><pre>6
30 10 60 10 60 50
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 3</h3><pre>40
</pre>

<p>If we follow the path <var>1</var> → <var>3</var> → <var>5</var> → <var>6</var>, the total cost incurred would be <var>|30 - 60| + |60 - 60| + |60 - 50| = 40</var>.</p></section>
</div>
