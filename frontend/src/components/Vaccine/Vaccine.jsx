import React, { Component, Fragment } from 'react';
import { Link } from 'react-router-dom';
import { Row, Col, ListGroup, ListGroupItem, ListGroupItemHeading, ListGroupItemText, UncontrolledCollapse, Button, CardBody, Card } from 'reactstrap';

import AddVaccine from './AddVacineModal';
import UpdateVaccine from './UpdateVacinneModal';

class Vaccine extends Component {

  state = { vaccine: null }
  
  constructor(props){
    super(props);
    this.state = {
      vaccines: [],
      isLoaded: false,

    }
  }

  componentDidMount() {
    fetch('https://jsonplaceholder.typicode.com/users')
    .then(res => res.json())
    .then(json => {
      this.setState({
        isLoaded: true,
        vaccine: json,
      })
    });
  }
  
  render() {

    var { isLoaded, vaccines } = this.state;

    if(!isLoaded) {
      return <div>Loading...</div>;
    }
    else {
      return (
        <Fragment>
        <div className="position-relative">
          <Row>
            <Col>
            <span className="pb-4 h2 text-dark border-bottom border-gray">All available Vaccines</span>
            </Col>
            <Col>
            <div className="align-right">
            <AddVaccine/>
            </div>
            </Col>
          </Row>
          
          <ListGroup>
        <ListGroupItem>
          <a href="" id="toggler" style={{ marginBottom: '1rem' }}>
            <ListGroupItemHeading>
              List group item heading
            </ListGroupItemHeading>
          </a>
        
          <ListGroupItemText>
          Donec id elit non mi porta gravida at eget metus. Maecenas sed diam eget risus varius blandit.
          </ListGroupItemText>
          <UncontrolledCollapse toggler="#toggler">
      <Card>
        <CardBody>
          <Row>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Nesciunt magni, voluptas debitis
          similique porro a molestias consequuntur earum odio officiis natus, amet hic, iste sed
          dignissimos esse fuga! Minus, alias.
          </Row>
          <Row className="align-right">
            <UpdateVaccine/>
            <Link to="/{id}/delete"><Button color="danger" style={{marginRight: 15}}>Delete</Button></Link>
          </Row>
        </CardBody>
      </Card>
    </UncontrolledCollapse>
        </ListGroupItem>
        <ListGroupItem>
          <ListGroupItemHeading>List group item heading</ListGroupItemHeading>
          <ListGroupItemText>
          Donec id elit non mi porta gravida at eget metus. Maecenas sed diam eget risus varius blandit.
          </ListGroupItemText>
        </ListGroupItem>
        <ListGroupItem>
          <ListGroupItemHeading>List group item heading</ListGroupItemHeading>
          <ListGroupItemText>
          Donec id elit non mi porta gravida at eget metus. Maecenas sed diam eget risus varius blandit.
          </ListGroupItemText>
        </ListGroupItem>
      </ListGroup>
          
        </div>
      </Fragment>
      );
    }
  }
  
}

export default Vaccine;