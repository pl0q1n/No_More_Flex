import scala.collection.mutable
import Main.AddTransaction
import com.twitter.io.Buf
import org.scalatest.FunSuite
import io.finch._

class MainTest extends FunSuite {
  test("Smoke post test") {
    Main.db = collection.mutable.Map.empty // DelayedInit semantics can be surprising (C)
    val initialSize = Main.db.size
    val transaction = Buf.Utf8(
      """
      {
        "sender": "foo",
        "receiver": "bar",
        "value": 42,
        "time": 1559481217
      }
      """
    )

    val request = Input.post("/add")
      .withBody[Application.Json](transaction)
    val response = Main.addTransaction(request).awaitOutputUnsafe()
    assert(response.map(_.value).contains(Main.Status(true)))
    assert(Main.db.size == initialSize + 1)
  }
  test("Smoke test of get transaction") {
    Main.db = mutable.Map(0 -> List(Transaction("a", "b", 42, 1559481217, None), 
                                    Transaction("master", "slave", 0, 1, None)))

    {
      val request = Input.get("/transactions", "sender" -> "a")
      val response = Main.getTransactions(request).awaitOutputUnsafe()
      assert(response.map(_.value).size == 1)
    }
    {
     val request = Input.get("/transactions")
     val response = Main.getTransactions(request).awaitOutputUnsafe()
     assert(response.map(_.value).head.size == 2)
    }
  }
}
