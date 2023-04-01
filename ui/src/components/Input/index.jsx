import Prop from 'prop-types';

export const Input = ({ onChange }) => {
  return <input onChange={onChange} type="text" />;
};

Input.propTypes = {
  onChange: Prop.func.isRequired,
};
