import { useState } from 'react';
import { Button } from '../../../components/Button';
import { Div } from '../../../components/Button/styles';
import { Input } from '../../../components/Input';

export const Create = () => {
  const [nameState, setNameState] = useState('');
  const [genderState, setGenderState] = useState('');
  const [ageState, setAgeState] = useState(0);

  const save = async () => {
    const userData = {
      name: nameState,
      gender: genderState,
      age: Number(ageState),
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

  return (
    <Div>
      Name: <Input onChange={updateName} />
      Gender: <Input onChange={updateGender} />
      Age: <Input onChange={updateAge} />
      <Button onClick={save}>Save</Button>
      <Button onClick={cancel}>Cancelar</Button>
    </Div>
  );
};
