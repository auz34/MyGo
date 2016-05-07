using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TreeSerDer
{
    class Program
    {
        static void Main(string[] args)
        {
            /*var node6 = new TreeNode(6);
            var node7 = new TreeNode(7);
            var node4 = new TreeNode(4) { left = node6, right = node7 };
            var node5 = new TreeNode(5);
            var node2 = new TreeNode(2) { left = node4, right = node5 };
            var node3 = new TreeNode(3);
            var node1 = new TreeNode(1) { left = node2, right = node3 };*/

            /*var node2 = new TreeNode(2);
            var node1 = new TreeNode(1) { right=node2 };*/
            var node0 = new TreeNode(0);

            var node1 = new TreeNode(1) { left = node0 };
            var node19 = new TreeNode(19);

            var node2 = new TreeNode(2) { left = node1 };
            var node18 = new TreeNode(18) { right = node19 };

            var node3 = new TreeNode(3) { left = node2 };
            var node17 = new TreeNode(17) { right = node18 };

            var node4 = new TreeNode(4) { left = node3 };
            var node16 = new TreeNode(16) { right = node17 };

            var node5 = new TreeNode(5) { left = node4 };
            var node15 = new TreeNode(15) { right = node16 };

            var node6 = new TreeNode(6) { left = node5 };
            var node14 = new TreeNode(14) { right = node15 };

            var node7 = new TreeNode(7) { left = node6 };
            var node13 = new TreeNode(13) { right = node14 };

            var node8 = new TreeNode(8) { left = node7 };
            var node12 = new TreeNode(12) { right = node13 };

            var node9 = new TreeNode(9) {left = node8};
            var node11 = new TreeNode(11) {right = node12};

            var node10 = new TreeNode(10) {left = node9, right = node11};

            var codec = new Codec();

            var s = codec.serialize(node10);
            Console.WriteLine(s);

            var newTree = codec.deserialize(s);

            Console.ReadLine();
        }
    }

   
    public class TreeNode {
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x) { val = x; }
    }

    public class Codec
    {

        // Encodes a tree to a single string.
        public string serialize(TreeNode root)
        {
            var queue = new Queue<TreeNode>();
            var result = new List<string>();

            queue.Enqueue(root);

            while (queue.Any())
            {
                var curNode = queue.Dequeue();
                if (curNode != null)
                {
                    result.Add(curNode.val.ToString());
                    queue.Enqueue(curNode.left);
                    queue.Enqueue(curNode.right);
                }
                else
                {
                    result.Add("null");
                }
            }

            return string.Join(",", result);
        }

        private TreeNode deserializeNode(string item)
        {
            if (item == "null")
            {
                return null;
            }

            return new TreeNode(int.Parse(item));
        }

        // Decodes your encoded data to tree.
        public TreeNode deserialize(string data)
        {
            if (string.IsNullOrEmpty(data))
            {
                return null;
            }

            var items = data.Split(',');

            var root = deserializeNode(items[0]);
            var parentNodes = new Queue<TreeNode>();
            parentNodes.Enqueue(root);
            var i = 1;
            while (i < items.Length)
            {
                var parent = parentNodes.Dequeue();
                if (parent != null)
                {
                    parent.left = this.deserializeNode(items[i++]);
                    if (i<items.Length)
                    {
                        parent.right = this.deserializeNode(items[i++]);
                    }

                    if (parent.left != null)
                    {
                        parentNodes.Enqueue(parent.left);
                    }

                    if (parent.right != null)
                    {
                        parentNodes.Enqueue(parent.right);
                    }
                }
                else
                {
                    i++;
                }
            }

            return root;
        }
    }

}
