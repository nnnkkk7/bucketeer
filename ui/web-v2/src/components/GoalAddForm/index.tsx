import { Dialog } from '@headlessui/react';
import { FC, memo } from 'react';
import { useFormContext } from 'react-hook-form';
import { useIntl } from 'react-intl';

import { messages } from '../../lang/messages';
import { Goal } from '../../proto/experiment/goal_pb';

export interface GoalAddFormProps {
  onSubmit: () => void;
  onCancel: () => void;
}

export const GoalAddForm: FC<GoalAddFormProps> = memo(
  ({ onSubmit, onCancel }) => {
    const { formatMessage: f } = useIntl();
    const methods = useFormContext();
    const {
      register,
      formState: { errors, isSubmitting, isValid },
      setValue,
      watch
    } = methods;

    const watchConnectionType = watch('connectionType');

    return (
      <div className="w-[500px]">
        <form className="flex flex-col">
          <div className="flex-1 h-0">
            <div className="py-6 px-4 bg-primary">
              <div className="flex items-center justify-between">
                <Dialog.Title className="text-lg font-medium text-white">
                  {f(messages.goal.add.header.title)}
                </Dialog.Title>
              </div>
              <div className="mt-1">
                <p className="text-sm text-indigo-300">
                  {f(messages.goal.add.header.description)}
                </p>
              </div>
            </div>
            <div
              className="
                flex-1
                flex flex-col
                justify-between
              "
            >
              <div
                className="
                  space-y-6 px-5 pt-6 pb-5
                  flex flex-col
                "
              >
                <div className="">
                  <label htmlFor="id">
                    <span className="input-label">{f({ id: 'id' })}</span>
                  </label>
                  <div className="mt-1">
                    <input
                      {...register('id')}
                      type="text"
                      name="id"
                      id="id"
                      className="input-text w-full"
                    />
                    <p className="input-error">
                      {errors.id && (
                        <span role="alert">{errors.id.message}</span>
                      )}
                    </p>
                  </div>
                </div>
                <div className="">
                  <label htmlFor="name">
                    <span className="input-label">{f({ id: 'name' })}</span>
                  </label>
                  <div className="mt-1">
                    <input
                      {...register('name')}
                      type="text"
                      name="name"
                      id="name"
                      className="input-text w-full"
                    />
                    <p className="input-error">
                      {errors.name && (
                        <span role="alert">{errors.name.message}</span>
                      )}
                    </p>
                  </div>
                </div>
                <div className="">
                  <label htmlFor="description" className="block">
                    <span className="input-label">
                      {f(messages.description)}
                    </span>
                    <span className="input-label-optional">
                      {' '}
                      {f(messages.input.optional)}
                    </span>
                  </label>
                  <div className="mt-1">
                    <textarea
                      {...register('description')}
                      id="description"
                      name="description"
                      rows={4}
                      className="input-text w-full"
                    />
                    <p className="input-error">
                      {errors.description && (
                        <span role="alert">{errors.description.message}</span>
                      )}
                    </p>
                  </div>
                </div>
                <div>
                  <label htmlFor="connections">
                    <span className="input-label">
                      {f(messages.goal.connection)}
                    </span>
                  </label>
                  <div className="mt-1">
                    {[
                      {
                        label: f(messages.goal.experiments),
                        value: Goal.ConnectionType.EXPERIMENT.toString()
                      },
                      {
                        label: f(messages.goal.autoOperations),
                        value: Goal.ConnectionType.OPERATION.toString()
                      }
                    ].map(({ label, value }) => (
                      <div
                        key={value}
                        className="flex items-center py-2 space-x-3"
                      >
                        <input
                          id={value}
                          type="radio"
                          value={value}
                          className="h-4 w-4 text-primary focus:ring-primary border-gray-300"
                          onChange={(e) =>
                            setValue('connectionType', e.target.value, {
                              shouldValidate: true
                            })
                          }
                          checked={watchConnectionType === value}
                        />
                        <label
                          htmlFor={value}
                          className="cursor-pointer text-sm text-gray-500"
                        >
                          {label}
                        </label>
                      </div>
                    ))}
                    <p className="input-error">
                      {errors.connectionType && (
                        <span role="alert">
                          {errors.connectionType.message}
                        </span>
                      )}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div className="flex-shrink-0 px-4 py-4 flex justify-end">
            <div className="mr-3">
              <button
                type="button"
                className="btn-cancel"
                disabled={false}
                onClick={onCancel}
              >
                {f(messages.button.cancel)}
              </button>
            </div>
            <button
              type="button"
              className="btn-submit"
              disabled={!isValid || isSubmitting}
              onClick={onSubmit}
            >
              {f(messages.button.submit)}
            </button>
          </div>
        </form>
      </div>
    );
  }
);
