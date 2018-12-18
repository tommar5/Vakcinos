import React, { Component, Fragment } from 'react';
import axios from 'axios';
import { ListGroup, ListGroupItem, ListGroupItemHeading, ListGroupItemText, UncontrolledCollapse, Button, CardBody, Card } from 'reactstrap';

class MyVaccine extends Component {

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
          
          <span className="d-block pb-4 h2 text-dark border-bottom border-gray">My Vaccines</span>
          
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
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Nesciunt magni, voluptas debitis
          similique porro a molestias consequuntur earum odio officiis natus, amet hic, iste sed
          dignissimos esse fuga! Minus, alias.
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

export default MyVaccine;