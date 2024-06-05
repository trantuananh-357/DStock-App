
fun main() {
    // Ví dụ sử dụng
   
    val xPoints : List<Double> = listOf(1.0, 2.0)
    val yPoints : List<Double> = listOf(5.0, 9.0)
    val n = xPoints.size

        var sumX : Double= 0.0
        var sumY :Double = 0.0
        var sumXY : Double= 0.0
        var sumXSquare : Double = 0.0

        for (i in 0 until n) {
            sumX += xPoints[i]
            sumY +=  yPoints[i]
            sumXY += xPoints[i] * yPoints[i]
            sumXSquare += xPoints[i] * xPoints[i]
        }

        val a = (n * sumXY - sumX * sumY) / (n * sumXSquare - sumX * sumX)
        val b = (sumY - a * sumX) / n
        println("a = " + a )
        println("b = " + b )


}