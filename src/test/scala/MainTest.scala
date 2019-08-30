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
        "time": "now"
      }
      """
    )

    val request = Input.post("/add")
      .withBody[Application.Json](transaction)
    val response = Main.addTransaction(request).awaitOutputUnsafe()
    assert(response.map(_.value).contains(Main.Status(true)))
    assert(Main.db.size == initialSize + 1)
  }
}
