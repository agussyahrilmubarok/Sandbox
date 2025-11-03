import PropTypes from "prop-types";
import { useCallback } from "react";
import { useDropzone } from "react-dropzone";
import { Center, Input, Image, Tooltip } from "@chakra-ui/react";

AvatarUploader.propTypes = {
  imageUrl: PropTypes.string.isRequired, // imageUrl should be a string and required
  onFieldChange: PropTypes.func.isRequired, // onFieldChange should be a function and required
  setFiles: PropTypes.func.isRequired, // setFiles should be a function and required
  onFileUpload: PropTypes.func, // onFileUpload function to upload file to the server
};

export function AvatarUploader({
  imageUrl,
  onFieldChange,
  setFiles,
  onFileUpload,
}) {
  const convertFileToUrl = (file) => URL.createObjectURL(file);

  const onDrop = useCallback(
    (acceptedFiles) => {
      setFiles(acceptedFiles);
      
      onFieldChange(convertFileToUrl(acceptedFiles[0])); // Change image preview

      
      // If onFileUpload is provided, call the function to upload the file to the server
      if (onFileUpload) {
        onFileUpload(acceptedFiles);
      }
    },
    [onFileUpload, setFiles, onFieldChange]
  );

  const { getRootProps, getInputProps } = useDropzone({
    onDrop,
    accept: {
      "image/jpeg": [],
      "image/png": [],
    },
  });

  return (
    <Center {...getRootProps()}>
      <Input {...getInputProps()} id="avatar" cursor="pointer" />
      <Tooltip label="Change your avatar">
        <Image
          alt="profile"
          rounded="full"
          h="24"
          w="24"
          objectFit="cover"
          cursor="pointer"
          mt="2"
          src={imageUrl}
        />
      </Tooltip>
    </Center>
  );
}
