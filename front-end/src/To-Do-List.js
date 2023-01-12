import React,{Component} from "react";
import axios from "axios";
import {Card, Header, Form, Input, Icon} from "semantic-ui-react"

let endpoint = "http://localhost:9000";

class ToDoList extends Component {
    constructor(props){
        super(props);

        this.state ={
            task:"",
            item:[],
        };
    }
    ComponentDidMount(){
        this.getTask();
    }
    onChange = (event) => {
        this.SetState({
            [event.target.name] : event.targe.value,
        })
    }

    onSubmit

    getTask = ()=> {
        axios.get(endpoint + ("/api/task").then((result)=> {
            if (res.data) {
                this.SetState({
                    items: res.data.map((item)=>{
                        let color = "yellow";
                        let style = {
                            wordWrap: "break-word",
                        };

                        if (item.status){
                            color="green";
                            style["textDecorationLine"] = "line-through";
                        }

                        return (
                           <Card key={item.id} color={color} fluid className="rough">
                                <Card.Content>
                                    <Card.Header textAlign="left">
                                        <div style={style}>{item.task}</div>
                                    </Card.Header>
                                        <Card.Meta textAlign="right">
                                            <Icon
                                            name="check circle"
                                            color="blue"
                                            onClick={() => this.updateTask(item.id)}
                                            />
                                            <span style={{paddingRight: 10}}>Undo</span>
                                            <Icon
                                            name="delete"
                                            color="red"
                                            onClick={()=> this.deleteTask(task.id)}
                                            />
                                            <span style={{paddingRight: 10}}>Delete</span>
                                        </Card.Meta> 
                                </Card.Content>
                           </Card> 
                        )
                    })
                })
            } else {
                this.SetState({
                    items:[],
                });
            }
        }));
    };
 
    updateTask = (id) => {
        axios.put(endpoint + "/api/task/" + id, {
            headers:{
                "Content-Type":"application/x-www-form-urlencoded",
            },
        }).then((res)=>{
            console.log(res);
            this.getTask();
        })
    }

    undoTask = (id) => {
        axios.put(endpoint + "api/task/" + id, {
            headers:{
                "Content-Type":"application/x-www-form-urlencoded",
            },
        }).then((res)=>{
            console.log(res)
            this.getTask();
        })
    }

    deleteTask = (id) => {
        axios.delete(endpoint + "api/deleteTask/" + id, {
            headers:{
                "Content-Type":"application/x-www-form-urlencoded",
            },
        }).then((res)=>{
            console.log(res)
            this.getTask();
        })
    }



    render(){
        return(
            <div>
                <div className="row">

                 <Header className="header" as="h2" color="yellow">
                    TO DO LIST
                 </Header>
                </div>
                <div className="row">
                    <Form onSubmit={this.onSubmit}>
                        <Input
                        type="text"
                        name="task"
                        onChange={this.onChange}
                        value={this.state.task}
                        fluid
                        placeholder="Create Task"
                        />
                    {/* {<Button> Create Task </Button>} */}
                    </Form>
                    </div>
                    <div className="row">
                        <Card.Group> {this.state.item} </Card.Group>
                    </div>
            </div>
        );
    }
}

export default ToDoList;