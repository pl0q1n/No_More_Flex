case class Transaction(sender: String,
                       receiver: String,
                       value: Int,
                       time: String,
                       category: Option[String])