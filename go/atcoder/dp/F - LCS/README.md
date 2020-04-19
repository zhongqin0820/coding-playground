<p>Score : <var>100</var> points</p>

<div class="part">
<section>
<h3>Problem Statement</h3><p>You are given strings <var>s</var> and <var>t</var>.
Find one longest string that is a subsequence of both <var>s</var> and <var>t</var>.</p>
</section>
</div>

<div class="part">
<section>
<h3>Notes</h3><p>A <em>subsequence</em> of a string <var>x</var> is the string obtained by removing zero or more characters from <var>x</var> and concatenating the remaining characters without changing the order.</p>
</section>
</div>

<div class="part">
<section>
<h3>Constraints</h3><ul>
<li><var>s</var> and <var>t</var> are strings consisting of lowercase English letters.</li>
<li><var>1 \leq |s|, |t| \leq 3000</var></li>
</ul>
</section>
</div>

<hr />
<div class="io-style">
<div class="part">
<section>
<h3>Input</h3><p>Input is given from Standard Input in the following format:</p>
<pre><var>s</var>
<var>t</var>
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Output</h3><p>Print one longest string that is a subsequence of both <var>s</var> and <var>t</var>.
If there are multiple such strings, any of them will be accepted.</p>
</section>
</div>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 1</h3><pre>axyb
abyxb
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 1</h3><pre>axb
</pre>

<p>The answer is <code>axb</code> or <code>ayb</code>; either will be accepted.</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 2</h3><pre>aa
xayaz
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 2</h3><pre>aa
</pre>

</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 3</h3><pre>a
z
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 3</h3><pre>
</pre>

<p>The answer is <code></code> (an empty string).</p>
</section>
</div>

<hr />
<div class="part">
<section>
<h3>Sample Input 4</h3><pre>abracadabra
avadakedavra
</pre>

</section>
</div>

<div class="part">
<section>
<h3>Sample Output 4</h3><pre>aaadara
</pre></section>
</div>
