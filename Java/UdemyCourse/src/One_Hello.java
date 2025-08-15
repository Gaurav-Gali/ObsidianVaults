public class One_Hello {
    public static void main(String[] args) {
        int[] arr = new int[5];
        for (int i = 0; i < arr.length; i++) {
            arr[i] = i+1;
        }

        for (int j : arr) {
            System.out.println(j);
        }
    }
}
