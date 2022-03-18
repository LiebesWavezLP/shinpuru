import styled from 'styled-components';
import { Guild, Member } from '../../lib/shinpuru-ts/src';
import { memberName } from '../../util/users';
import { Container } from '../Container';
import { DiscordImage } from '../DiscordImage';
import { Embed } from '../Embed';
import { Flex } from '../Flex';
import { RoleList } from '../RoleList';
import { Clickable } from '../styleParts';

interface Props {
  member?: Member;
  guild?: Guild;
}

const StyledContainer = styled(Container)`
  ${Clickable(1.01)}

  display: flex;
  width: 100%;

  > img {
    height: 4em;
    margin-right: 1em;
  }

  > div {
    > * {
      margin: 0 0 0.6em 0;
      &:last-child {
        margin: 0;
      }
    }
  }

  p,
  h2 {
    margin: 0;
  }
`;

const Header = styled(Flex)`
  @media screen and (max-width: 800px) {
    > h2 {
      display: block;
      width: 100%;
    }
  }

  flex-wrap: wrap;

  > * {
    margin-right: 0.5em !important;
    &:last-child {
      margin-right: 0 !important;
    }
  }
`;

export const MemberLarge: React.FC<Props> = ({ member, guild }) => {
  return !!member && !!guild ? (
    <StyledContainer>
      <DiscordImage src={member.avatar_url} />
      <div>
        <Header>
          <h2>{memberName(member, false)}</h2>
          <small>
            {member.user.username}#{member.user.discriminator}
          </small>
          <Embed>{member.user.id}</Embed>
        </Header>
        <RoleList guildroles={guild.roles!} roleids={member.roles} />
      </div>
    </StyledContainer>
  ) : (
    <>loading...</>
  );
};
