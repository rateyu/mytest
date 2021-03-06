# 二分查找（上）

## 算法描述

二分查找（Binary Search）也叫折半查找，是针对有序数据集合的查找算法。其描述十分简单：

* 折半取中，判断元素与目标元素的大小关系
    * 小于——往前继续折半
    * 大于——往后继续折半
    * 等于——返回

关于它的复杂度分析，参见[谈谈基于比较的排序算法的复杂度下界](https://liam.page/2018/08/28/lower-bound-of-comparation-based-sort-algorithm/)中的相关信息。它的复杂度是 $O(\log n)$。

## $O(\log n)$ 的惊人之处

在 42 亿个数据中用二分查找一个数据，最多需要比较 32 次。

## 适用场景

* 依赖顺序表结构
* 数据本身必须有序
* 数据量相对比较元素的开销要足够大——不然遍历即可
* 数据量相对内存空间不能太大——不然顺序表装不下
