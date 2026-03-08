import PropTypes from 'prop-types';
import {
  Button,
  AlertDialog,
  AlertDialogBody,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogContent,
  AlertDialogOverlay,
} from "@chakra-ui/react";

DeleteConfirmation.propTypes = {
  alertTitle: PropTypes.string.isRequired,  // Ensures alertTitle is a string and required
  handleClick: PropTypes.func.isRequired,   // Ensures handleClick is a function and required
  isOpen: PropTypes.bool.isRequired,        // Ensures isOpen is a boolean and required
  onClose: PropTypes.func.isRequired,       // Ensures onClose is a function and required
};

export default function DeleteConfirmation({
  alertTitle,
  handleClick,
  isOpen,
  onClose,
}) {
  return (
    <AlertDialog isOpen={isOpen} onClose={onClose}>
      <AlertDialogOverlay>
        <AlertDialogContent>
          <AlertDialogHeader fontSize="lg" fontWeight="bold">
            {alertTitle}
          </AlertDialogHeader>

          <AlertDialogBody>
            Are you sure? You can`t undo this action.
          </AlertDialogBody>

          <AlertDialogFooter>
            <Button onClick={onClose}>Cancel</Button>
            <Button colorScheme="red" onClick={handleClick} ml={3}>
              Delete
            </Button>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialogOverlay>
    </AlertDialog>
  );
}
