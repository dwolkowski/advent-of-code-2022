# Task Content
 - All tasks can be found on **[Advent of Code 2022](https://adventofcode.com/2022)** 
 
## Part One

**Content**
> The Elves begin to set up camp on the beach. To decide whose tent gets  
> to be closest to the snack storage, a giant Rock Paper Scissors tournament  
> is already in progress.

> Appreciative of your help yesterday, one Elf gives you an encrypted strategy  
> guide (your puzzle input) that they say will be sure to help you win.  
> "The first column is what your opponent is going to play: A for Rock, B for  
> Paper, and C for Scissors. The second column--" Suddenly, the Elf is called  
> away to help with someone's tent.  

> The second column, you reason, must be what you should play  
> in response: X for Rock, Y for Paper, and Z for Scissors.  
> Winning every time would be suspicious, so the responses  
> must have been carefully chosen.

> The winner of the whole tournament is the player with the highest score.  
> Your total score is the sum of your scores for each round. The score for  
> a single round is the score for the shape you selected (1 for Rock, 2 for  
> Paper, and 3 for Scissors) plus the score for the outcome of the  
> round (0 if you lost, 3 if the round was a draw, and 6 if you won).

> What would your total score be if everything goes exactly  
> according to your strategy guide?

<details>
  <summary><strong> Example </strong></summary>

Suppose you were given the following strategy guide:

<pre>A Y
B X
C Z</pre>

This strategy guide predicts and recommends the following:

 - In the first round, your opponent will choose `Rock (A)`, and you should choose `Paper (Y)`. This ends in a win for you with a score of **8** (2 because you chose Paper + 6 because you won).
 - In the second round, your opponent will choose `Paper (B)`, and you should choose `Rock (X)`. This ends in a loss for you with a score of **1** (1 + 0).
 - The third round is a draw with both players choosing `Scissors (C and Z)`, giving you a score of 3 + 3 = **6**.

<br />

In this example, if you were to follow the strategy guide, you would get a total score of **`15`** (8 + 1 + 6).

</details>

<details>
  <summary><strong> Answer </strong></summary>
  <div align="right">
    
   My puzzle answer was: **`15572`**
    
  </div>
</details> 

___

## Part Two

**Content**
> The Elf finishes helping with the tent and sneaks back over to you.  
> "Anyway, the second column says how the round needs to end: X means  
> you need to lose, Y means you need to end the round in a draw, and  
> Z means you need to win. Good luck!"   

> The total score is still calculated in the same way, but now  
> you need to figure out what shape to choose so the round ends as indicated. 

> Following the Elf's instructions for the second column, what would your  
> total score be if everything goes exactly according to your strategy guide?


<details>
  <summary><strong> Example </strong></summary>

The example above now goes like this:

 - In the first round, your opponent will choose `Rock (A)`, and you need the round to end in a draw (Y), so you also choose Rock. This gives you a score of 1 + 3 = **4**.
 - In the second round, your opponent will choose `Paper (B)`, and you choose Rock so you lose (X) with a score of 1 + 0 = **1**.
 - In the third round, you will defeat your opponent's `Scissors` with `Rock` for a score of 1 + 6 = **7**.

<br />

Now that you're correctly decrypting the ultra top secret strategy guide, you would get a total score of **12**.

</details>

<details>
  <summary><strong> Answer </strong></summary>
  <div align="right">
    
   My puzzle answer was: **`16098`**
    
  </div>
</details> 

___

<br />
