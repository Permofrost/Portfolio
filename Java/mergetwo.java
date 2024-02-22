public class mergetwo {
    // Define the Node class
    static class Node {
        int val;
        Node next;

        Node(int val) {
            this.val = val;
        }
    }

    public Node merging(Node a, Node b) {
        Node head = new Node(0);
        Node handler = head;

        while(a != null && b != null) {
            if(a.val < b.val) {
                handler.next = a;
                a = a.next;
            } else {
                handler.next = b;
                b = b.next;
            }
            handler = handler.next;
        }

        if(a != null) {
            handler.next = a;
        } else if(b != null) {
            handler.next = b;
        }

        return head.next;
    }

    public static void main(String[] args) {
        mergetwo morg = new mergetwo();

        // Create the first list
        Node l1 = new Node(1);
        l1.next = new Node(4);
        l1.next.next = new Node(7);

        // Create the second list
        Node l2 = new Node(2);
        l2.next = new Node(3);
        l2.next.next = new Node(5);
        l2.next.next.next = new Node(6);

        // Merge the lists
        Node result = morg.merging(l1, l2);

        // Print the result
        while (result != null) {
            System.out.print(result.val + " ");
            result = result.next;
        }
    }
}