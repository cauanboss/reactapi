import Prop from 'prop-types';

export const Button = ({ children, onClick }) => {
  return <button onClick={onClick}>{children}</button>;
};

Button.propTypes = {
  children: Prop.node.isRequired,
  onClick: Prop.func.isRequired,
};
