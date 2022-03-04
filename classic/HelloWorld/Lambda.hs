---
--- Just for Fun
---
--- Author:       vac
--- Organization: STEM
--- LICENCE:      MIT
--- Copyright:    2014-2018

module Main (main) where

-- function declaration
sayHello :: IO ()
sayHello =
    putStrLn "hello, world" -- output some text

-- entry point of the program
main :: IO ()
main =
    sayHello -- function invocation
