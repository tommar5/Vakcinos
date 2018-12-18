import React, { Component } from 'react';
import { Button, Modal, ModalHeader, ModalBody, ModalFooter, Form, FormGroup, Label, Input } from 'reactstrap';

class UpdateVaccineModal extends Component {
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
            <Button color="warning" onClick={this.toggle} style={{marginRight: 15}}>Update information</Button>{' '}
            <Modal isOpen={this.state.modal} toggle={this.toggle} className={this.props.className}>
                <Form>
                <ModalHeader toggle={this.toggle}>Update vaccine</ModalHeader>
                    <ModalBody>
                        <FormGroup>
                            <Label for="title">Title</Label>
                            <Input type="text" name="title" id="title" />
                        </FormGroup>
                        <FormGroup>
                            <Label for="description">Description</Label>
                            <Input type="text" name="description" id="description" />
                        </FormGroup> 
                        <FormGroup>
                            <Label for="lot">LOT</Label>
                            <Input type="text" name="lot" id="lot" />
                        </FormGroup>
                        <FormGroup>
                            <Label for="expiryDate">Expiry Date</Label>
                            <Input type="date" name="expiryDate" id="expiryDate" />
                        </FormGroup>   
                        <FormGroup>
                            <Label for="fromAge">Vaccinating from age</Label>
                            <Input type="text" name="fromAge" id="fromAge" />
                        </FormGroup> 
                        <FormGroup>
                            <Label for="cost">Cost</Label>
                            <Input type="text" name="cost" id="cost" />
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

export default UpdateVaccineModal;