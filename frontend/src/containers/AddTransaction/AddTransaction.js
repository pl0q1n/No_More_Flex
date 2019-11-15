import React, {Component} from 'react'
import Button from '../../components/Button/Button'
import axios from 'axios';

class AddTransaction extends Component {
    constructor(props) {
        super(props);
        this.state = {
        };
    
        this.handleInputChange = this.handleInputChange.bind(this);
        this.sendTransaction = this.sendTransaction.bind(this);
      }
    
      handleInputChange(event) {
        const target = event.target;
        const value = target.value;
        const name = target.name;
    
        this.setState({
          [name]: value
        });
      }
    
      sendTransaction() {    
        console.log("Sending transaction: " + JSON.stringify(this.state));
        axios({
            method: 'post',
            url: 'http://localhost:8080/transactions/add',
            data: this.state
        })
        
        
      }
    
      render() {
        return (
          <form>
            <label>
              Sender: &nbsp;
              <input
                name="sender"
                type="text"
                value={this.state.sender}
                onChange={this.handleInputChange} />
                <br />
            </label>
            <label>
              Receiver: &nbsp;
              <input
                name="receiver"
                type="text"
                value={this.state.receiver}
                onChange={this.handleInputChange} />
                <br />
            </label>
            <label>
              Value: &nbsp;
              <input
                name="value"
                type="number"
                value={this.state.value}
                onChange={this.handleInputChange} />
                <br />
            </label>
            <label>
              Timestamp: &nbsp;
              <input
                name="timestamp"
                type="number"
                value={this.state.timestamp}
                onChange={this.handleInputChange} />
                <br />
            </label>
            <label>
              Category: &nbsp;
              <input
                name="category"
                type="text"
                value={this.state.category}
                onChange={this.handleInputChange} />
                <br />
            </label>
            
          <Button label={"Send transaction"} handleClick={this.sendTransaction}/>
          </form>
        );
      }
}

export default AddTransaction