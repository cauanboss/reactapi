/* eslint-disable react/prop-types */
import { useState } from 'react';
import { Button } from '../../../components/Button';
import { Div } from '../../../components/Button/styles';
import { Input } from '../../../components/Input';

// eslint-disable-next-line react/prop-types
export const Create = ({ defaultValues }) => {
  const [nameState, setNameState] = useState('');
  const [genderState, setGenderState] = useState('');
  const [ageState, setAgeState] = useState(0);
  const [usernameState, setUsernameState] = useState(0);
  const [passwordState, setPasswordState] = useState(0);

  const save = async () => {
    const userData = {
      name: nameState,
      gender: genderState,
      age: Number(ageState),
      password: passwordState,
      username: usernameState,
    };

    try {
      const user = await fetch('http://localhost/api/v1/user', {
        body: JSON.stringify(userData),
        method: 'post',
        mode: 'no-cors',
      });

      console.log(user);
      console.log('save');
    } catch (error) {
      console.log(error);
    }
  };
  const cancel = () => {
    console.log('cancel');
  };
  const updateName = (e) => {
    setNameState(e.target.value);
  };
  const updateGender = (e) => {
    setGenderState(e.target.value);
  };
  const updateAge = (e) => {
    setAgeState(e.target.value);
  };
  const updateUsername = (e) => {
    setUsernameState(e.target.value);
  };
  const updatePassword = (e) => {
    setPasswordState(e.target.value);
  };

  return (
    <Div>
      Name: <Input value={defaultValues?.name || ''} onChange={updateName} />
      Gender: <Input value={defaultValues?.gender || ''} onChange={updateGender} />
      Age: <Input value={defaultValues?.age || ''} onChange={updateAge} />
      Age: <Input value={defaultValues?.username || ''} onChange={updateUsername} />
      Age: <Input value={defaultValues?.password || ''} onChange={updatePassword} />
      <Button onClick={save}>Save</Button>
      <Button onClick={cancel}>Cancelar</Button>
    </Div>
  );
};
