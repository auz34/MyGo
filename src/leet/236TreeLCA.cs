using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LCA
{
    class Program
    {
        static void Main(string[] args)
        {
            var node7 = new TreeNode(7);
            var node4 = new TreeNode(4);

            var node6 = new TreeNode(6);
            var node2 = new TreeNode(2) {left = node7, right = node4};
            var node0 = new TreeNode(0);
            var node8 = new TreeNode(8);

            var node5 = new TreeNode(5) {left=node6, right = node2};
            var node1 = new TreeNode(1) {left = node0, right = node8};

            var node3 = new TreeNode(3) {left = node5, right = node1};

            var s = new Solution();

            Console.WriteLine(s.LowestCommonAncestor(node3, node4, node6).val);
            Console.WriteLine(s.LowestCommonAncestor(node3, node8, node6).val);
            Console.WriteLine(s.LowestCommonAncestor(node3, node5, node1).val);
            Console.WriteLine(s.LowestCommonAncestor(node3, node4, node5).val);

            Console.ReadLine();

        }
    }

    public class TreeNode
    {
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x) { val = x; }
    }

    public class Solution
    {
        private class TraverseResult
        {
            public bool containsFirst;

            public bool containsSecond;

            public TreeNode lca;

            public bool IsReady()
            {
                return this.lca != null;
            }

            public void AddNode(TreeNode node, TreeNode first, TreeNode second)
            {
                this.containsFirst = this.containsFirst || (node == first);
                this.containsSecond = this.containsSecond || (node == second);

                if (this.lca == null && this.containsFirst && this.containsSecond)
                {
                    this.lca = node;
                }
            }

            public TraverseResult Merge(TreeNode node, TraverseResult a)
            {
                if (this.IsReady())
                {
                    return this;
                }

                if (a.IsReady())
                {
                    return a;
                }

                var result = new TraverseResult()
                                 {
                                     containsFirst = this.containsFirst || a.containsFirst,
                                     containsSecond = this.containsSecond || a.containsSecond
                                 };

                if (result.containsFirst && result.containsSecond)
                {
                    result.lca = node;
                }

                return result;
            }
        }

        private TraverseResult LCASearch(TreeNode node, TreeNode first, TreeNode second)
        {
            if (node == null)
            {
                return new TraverseResult() { containsFirst = false, containsSecond = false, lca = null};
            }

            var left = this.LCASearch(node.left, first, second);
            left.AddNode(node, first, second);
            if (left.IsReady())
            {
                return left;
            }

            var right = this.LCASearch(node.right, first, second);
            return left.Merge(node, right);
        }

        public TreeNode LowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q)
        {
            return this.LCASearch(root, p, q).lca;
        }
    }
}
