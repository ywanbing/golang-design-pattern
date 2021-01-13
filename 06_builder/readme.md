# 生成器模式

定义：
  > 一种创建型设计模式，使你能够分步骤创建复杂对象。 该模式允许你使用相同的创建代码生成不同类型和形式的对象。  
  > 生成器模式（Builder）是使用多个“小型”工厂来最终创建出一个完整对象。  
  > ![n](https://refactoringguru.cn/images/patterns/content/builder/builder-zh.png)

比如：
  > igloo­Builder冰屋生成器与 normal­Builder普通房屋生成器可建造不同类型房屋， 即 igloo冰屋和 normal­House普通房屋 。 每种房屋类型的建造步骤都是相同的

本质： 分离整体构建算法和部件构造。    

何时使用： 
  * 如果创建对象的算法，应该独立于该对象的组成部分以及它们的装配方式时。  
  * 如果同一个构建过程有着不同的表现时。