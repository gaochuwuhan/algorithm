<div>
	在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
</div>
<div>
	[
</div>
<div>
	[1,2,8,9],<br />
[2,4,9,12],<br />
[4,7,10,13],<br />
[6,8,11,15]<br />
</div>
<div>
	]
</div>
<div>
	<p style="color:#262626;">
		给定 target&nbsp;= 7，返回&nbsp;true。
	</p>
	<div style="color:#262626;">
		给定&nbsp;target&nbsp;=&nbsp;3，返回&nbsp;false。
	</div>
	<div style="color:#262626;">
		<br />
	</div>
	<div style="color:#262626;">
		数据范围：矩阵的长宽满足 <img src="https://www.nowcoder.com/equation?tex=0%20%5Cle%20n%2Cm%20%5Cle%20%20500" /> ， 矩阵中的值满足 <img src="https://www.nowcoder.com/equation?tex=-10%5E9%20%5Cle%20val%20%5Cle%2010%5E9" alt="-10^9 \le val \le 10^9" /><br />
进阶：空间复杂度 <img src="https://www.nowcoder.com/equation?tex=O(1)" /> ，时间复杂度 <img src="https://www.nowcoder.com/equation?tex=O(n%2Bm)" /><br />
	</div>
</div><div><br></div><div><div>Related Topics</div><div><li>数组</li></div></div><br>示例:<br>输入:7,[[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]<br>输出:true<br>