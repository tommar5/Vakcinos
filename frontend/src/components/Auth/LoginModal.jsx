import React, { Component } from 'react';
import { Button, Modal, ModalHeader, ModalBody, ModalFooter, Form, FormGroup, Label, Input } from 'reactstrap';

class LoginModal extends Component {
    constructor(props) {
        super(props);
        this.state = {
          modal: false
        };
    
        this.toggle = this.toggle.bind(this);
      }
    
      toggle() {
        this.setState({
          modal: !this.state.modal
        });
      }
    
      render() {
        return (
          <div>
            <Button color="info" outline onClick={this.toggle} style={{marginRight: 15}}>Login</Button>{' '}
            <Modal isOpen={this.state.modal} toggle={this.toggle} className={this.props.className}>
                <Form>
                    <ModalHeader toggle={this.toggle}>Login to account</ModalHeader>
                    <ModalBody>
              
                        <FormGroup>
                            <Label for="exampleEmail">Email</Label>
                            <Input type="email" name="email" id="exampleEmail" placeholder="Enter your email" />
                        </FormGroup>
                        <FormGroup>
                            <Label for="examplePassword">Password</Label>
                            <Input type="password" name="password" id="examplePassword" placeholder="Enter your password" />
                        </FormGroup>  
                    </ModalBody>
                    <ModalFooter>
                        <Button color="primary" type="submit" onClick={this.toggle}>Login</Button>{' '}
                        <Button color="secondary" onClick={this.toggle}>Cancel</Button>
                    </ModalFooter>
                </Form>
            </Modal>
          </div>
        );
      }
}

export default LoginModal;