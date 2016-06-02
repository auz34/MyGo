using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace RegexMatching
{
    using System.Globalization;
    using System.Xml.Schema;

    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();

            Console.WriteLine(s.IsMatch(string.Empty, string.Empty));
            Console.WriteLine(s.IsMatch(string.Empty, ".*"));
            Console.WriteLine(s.IsMatch("asddaefdsaag", ".*"));
            Console.WriteLine(s.IsMatch("asddaefdsaag", ".*b"));
            Console.WriteLine(s.IsMatch("aa", "a"));
            Console.WriteLine(s.IsMatch("aa", "aa"));
            Console.WriteLine(s.IsMatch("aaa", "aa"));
            Console.WriteLine(s.IsMatch("aaa", "a*"));
            Console.WriteLine(s.IsMatch("ab", ".*"));
            Console.WriteLine(s.IsMatch("aab", "c*a*b"));
            Console.WriteLine(s.IsMatch("aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*a*a*b"));

            /*;
            Console.WriteLine(s.IsMatch(string.Empty, "."));
            Console.WriteLine(s.IsMatch(string.Empty, "*."));
            Console.WriteLine(s.IsMatch("c", "*."));
            Console.WriteLine(s.IsMatch("csaddsdssd", "*."));
            Console.WriteLine(s.IsMatch("abccd", "abc."));*/
            Console.ReadLine();
        }
    }

    public class Solution
    {
        private class Pattern
        {
            public List<string> Tokens = new List<string>();
            public Pattern(string p)
            {
                var i = 0;
                while (i < p.Length)
                {
                    var lastToken = this.Tokens.LastOrDefault();
                    if (i < p.Length - 1 && p[i+1] == '*')
                    {
                        var newToken = p.Substring(i, 2);
                        if (newToken != ".*")
                        {
                            if (lastToken != ".*" && lastToken != newToken)
                            {
                                this.Tokens.Add(newToken);
                            }
                        }
                        else
                        {
                            while (!string.IsNullOrEmpty(lastToken) && lastToken.Length == 2)
                            {
                                this.Tokens.RemoveAt(this.Tokens.Count - 1);
                                lastToken = this.Tokens.LastOrDefault();
                            }

                            this.Tokens.Add(newToken);
                        }
                        
                        i += 2;
                        continue;
                    }

                    this.Tokens.Add(p.Substring(i, 1));
                    i++;
                }
            }

            public int Length
            {
                get
                {
                    return this.Tokens.Count;
                }
            } 
        }

        private class State
        {
            public Pattern Pattern { get; set; }

            public int PatternIndex { get; set; }

            private bool IsStarToken(string token)
            {
                return token.Length == 2;
            }

            private bool MatchStarToken(string token, char ch)
            {
                return token[0] == '.' || token[0] == ch;
            }

            private List<State> GetNextStatesForStartToken(char ch)
            {
                var result = new List<State>();
                var token = this.Pattern.Tokens[this.PatternIndex + 1];
                if (this.MatchStarToken(token, ch))
                {
                    result.Add(this);
                }

                var next = new State() { Pattern = this.Pattern, PatternIndex = this.PatternIndex + 1 };
                result.AddRange(next.GetNextStates(ch));
                return result;
            } 


            public List<State> GetNextStates(char ch)
            {
                if (this.PatternIndex + 1 >= this.Pattern.Length)
                {
                    return new List<State>();
                }

                var token = this.Pattern.Tokens[this.PatternIndex + 1];
                if (this.IsStarToken(token))
                {
                    return GetNextStatesForStartToken(ch);
                }

                var result = new List<State>();
                
                if (token == "." || token[0] == ch)
                {
                    result.Add(new State { Pattern = this.Pattern, PatternIndex = this.PatternIndex + 1 });
                }

                return result;
            }

            public bool IsValid()
            {
                if (PatternIndex + 1 == Pattern.Length)
                {
                    return true;
                }

                for (var i = PatternIndex + 1; i < Pattern.Length; i++)
                {
                    if (!IsStarToken(Pattern.Tokens[i]))
                    {
                        return false;
                    }
                }

                return true;
            }
        }
        public bool IsMatch(string s, string p)
        {
            var curStates = new List<State> { new State { Pattern = new Pattern(p), PatternIndex = -1 }};

            foreach (var ch in s)
            {
                var newStates = new List<State>();
                foreach (var state in curStates)
                {
                    newStates.AddRange(state.GetNextStates(ch));
                }

                curStates = newStates;
            }

            return curStates.Any(st => st.IsValid());
        }
    }
}
