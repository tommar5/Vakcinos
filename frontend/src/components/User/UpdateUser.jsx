import React, { Component } from 'react';
import { Button, Modal, ModalHeader, ModalBody, ModalFooter, Form, FormGroup, Label, Input } from 'reactstrap';

class UpdateUser extends Component {
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
          <div className="div-inline">
            <Button className="font-weight-bold" color="warning" onClick={this.toggle} style={{marginRight: 15}}>Edit</Button>{' '}
            <Modal isOpen={this.state.modal} toggle={this.toggle} className={this.props.className}>
                <Form>
                    <ModalHeader toggle={this.toggle}>Edit User information</ModalHeader>
                    <ModalBody>
              
                    <FormGroup>
                            <Label for="regName">Name</Label>
                            <Input type="text" name="name" id="regName" />
                        </FormGroup>
                        <FormGroup>
                            <Label for="regSurname">Surname</Label>
                            <Input type="text" name="surname" id="regSurname" />
                        </FormGroup>      
                        <FormGroup>
                            <Label for="regEmail">Email</Label>
                            <Input type="email" name="email" id="regEmail" />
                        </FormGroup>
                        <FormGroup>
                            <Label for="regBirthday">Birthday</Label>
                            <Input type="date" name="birthday" id="regBirthday" />
                        </FormGroup>   
                        <FormGroup>
                            <Label for="regPhone">Phone number</Label>
                            <Input type="text" name="phone" id="regPhone" />
                        </FormGroup> 
                        <FormGroup>
                            <Label for="regCity">City</Label>
                            <Input type="text" name="city" id="regCity" />
                        </FormGroup>
                        <FormGroup>
                            <Label for="regAddress">Address</Label>
                            <Input type="text" name="address" id="regAddress" />
                        </FormGroup>  
                    </ModalBody>
                    <ModalFooter>
                        <Button color="primary" type="submit" onClick={this.toggle}>Update</Button>{' '}
                        <Button color="secondary" onClick={this.toggle}>Cancel</Button>
                    </ModalFooter>
                </Form>
            </Modal>
          </div>
        );
      }
}

export default UpdateUser;