import React, { FC, memo } from 'react';
import ReactCreatableSelect from 'react-select/creatable';

export enum FormatCreateLabel {
  NONE = 'none'
}

export interface Option {
  value: string;
  label: string;
}

export interface CreatableSelectProps {
  options?: Option[];
  disabled?: boolean;
  isSearchable?: boolean;
  defaultValues?: Option[];
  closeMenuOnSelect?: boolean;
  className?: string;
  onChange: (options: Option[]) => void;
  value?: Option;
  placeholder?: string;
  formatCreateLabel?: FormatCreateLabel;
}

const textColor = '#3F3F46';
const textColorDisabled = '#6B7280';
const backgroundColor = 'white';
const backgroundColorDisabled = '#F3F4F6';
const borderColor = '#D1D5DB';
const fontSize = '0.875rem';
const lineHeight = '1.25rem';
const minHeight = '2.5rem';

export const colourStyles = {
  control: (styles, { isDisabled }) => ({
    ...styles,
    backgroundColor: isDisabled ? backgroundColorDisabled : backgroundColor,
    borderColor: borderColor,
    '&:hover': {
      borderColor: borderColor
    },
    fontSize: fontSize,
    lineHeight: lineHeight,
    minHeight: minHeight,
    '*': {
      boxShadow: 'none !important'
    }
  }),
  option: (styles, { isFocused }) => {
    return {
      ...styles,
      backgroundColor: isFocused ? backgroundColor : null,
      color: textColor
    };
  },
  menu: (base) => ({
    ...base,
    fontSize: fontSize,
    lineHeight: lineHeight,
    color: textColor
  }),
  multiValueLabel: (base, { isDisabled }) => ({
    ...base,
    color: isDisabled ? textColorDisabled : textColor
  }),
  singleValue: (styles, { isDisabled }) => {
    return {
      ...styles,
      color: isDisabled ? textColorDisabled : textColor
    };
  }
};

export const CreatableSelect: FC<CreatableSelectProps> = memo(
  ({
    disabled,
    isSearchable,
    className,
    onChange,
    options,
    defaultValues,
    closeMenuOnSelect,
    value,
    placeholder = '',
    formatCreateLabel
  }) => {
    return (
      <ReactCreatableSelect
        isMulti
        options={options}
        placeholder={placeholder}
        className={className}
        classNamePrefix="react-select"
        styles={colourStyles}
        components={{
          DropdownIndicator: null
        }}
        isDisabled={disabled}
        isSearchable={isSearchable}
        value={value}
        defaultValue={defaultValues}
        onChange={onChange}
        closeMenuOnSelect={closeMenuOnSelect}
        formatCreateLabel={(inputValue) =>
          formatCreateLabel === FormatCreateLabel.NONE
            ? `"${inputValue}"`
            : `Create "${inputValue}"`
        }
      />
    );
  }
);
