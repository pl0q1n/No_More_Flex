import org.scalatest.FunSuite
import io.finch._
import com.twitter.io.Buf

class MainTest extends FunSuite {
  test("Smoke post test") {
    val inp = Input.post("/value").withBody[Application.Json](Buf.Utf8("""{"name":"foo","age":42}"""))
    val actual_msg = Main.processPost(inp).awaitOutputUnsafe().map(_.value)
    val expected_msg = Some(Main.Message("post Value"))
    assert(actual_msg == expected_msg)
  }
}
