import React, { useState } from 'react'
import { Link } from 'react-router-dom'
import '../MyStyles.css'
import {Password} from 'primereact/password';
import { Button } from 'primereact/button';
import { InputText } from 'primereact/inputtext';
import { register } from '../api/Login.js';
import { fileUpload } from '../api/file.js';
import { Card } from 'primereact/card';
import { FileUpload } from 'primereact/fileupload';
import 'primeflex/primeflex.css';
 
const SignUp = ({history}) => 
{

    const [user, setUser] = useState({
        username: '',
        name: '',
        dp: '',
        bio: '',
        password: '',
    });

    const [image, setImage] = useState(null);
    
    const [err, setErr] = useState('');

    const handleChange = e => {
            const { name, value } = e.target;
            setUser(prevState => ({
                ...prevState,
                [name]: value
            }));
        };

    const onImageUpload = e =>{
        setImage(e.files[0])
    }

    const registerUser = async () => {
        try{
            let fileId = await fileUpload(image)
            setUser(prevState =>({
                ...prevState,
                dp: fileId
            }));
            console.log("File uploaded id: " + fileId);
            console.log("User: " + user.dp);

            let res = await register({
                ...user,
                dp: fileId
            })
            history.push('/project')
        }
        catch(err){
            console.log(err)
            setErr(err);
        }
        //history.push('/tasks')
    }

    return (
        <div className="signup-page p-grid p-justify-center p-align-center">
            <div className="p-col-4">
                <Card className="signup-card">
                    <h2> Register </h2>
                    <h5>NAME</h5>
                    <InputText name="name" value={user.name} onChange={handleChange}/>
                    <h5>ORGANIZATION</h5>
                    <InputText name="bio" value={user.bio} onChange={handleChange}/>
                    <h5>USERNAME</h5>
                    <InputText name="username" value={user.username} onChange={handleChange}/>
                    <h5>PASSWORD</h5>
                    <Password name="password" value={user.password} onChange={handleChange}/>
                    <h5>Avatar</h5>
                    <FileUpload mode="basic" name="image" url='http://localhost:8080/api/upload-file/' accept="image/*" maxFileSize={5000000} onSelect={onImageUpload} />

                    <br/>
                    <div className="err">{err} </div>
                    <Button label="Sign Up" onClick={() => 
                        registerUser()
                    } />
                    <br/>
                    <br/>
                    
                    <h5> 
                        Already have an account?
                    </h5>
                    <Link to="/">
                        Login
                    </Link>
                </Card>
            </div>
        </div>
    )
}

export default SignUp;